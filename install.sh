#!/bin/bash
set -e

resolve_latest_version() {
    local tags testing_tag latest_tag
    tags=$(curl -fsSL "https://api.github.com/repos/gilangji/agy-mobile/releases?per_page=100" 2>/dev/null | grep '"tag_name":' | sed -E 's/.*"tag_name": "([^"]+)".*/\1/' || true)
    if [ -n "$tags" ]; then
        latest_tag=$(printf '%s\n' "$tags" | grep -E '^v[0-9]+(\.[0-9]+)*$' | sort -V | tail -n 1 || true)
        if [ -n "$latest_tag" ]; then
            echo "$latest_tag"
            return
        fi
        testing_tag=$(printf '%s\n' "$tags" | grep -E '^v[0-9]+\.[0-9]+\.testing\.[0-9]+$' | sort -V | tail -n 1 || true)
        if [ -n "$testing_tag" ]; then
            echo "$testing_tag"
            return
        fi
        printf '%s\n' "$tags" | head -n 1
        return
    fi
    echo "v1.1.0"
}

REQUESTED_VERSION="${1:-${VERSION:-}}"
if [ -n "$REQUESTED_VERSION" ] && [ "$REQUESTED_VERSION" != "latest" ]; then
    VERSION="$REQUESTED_VERSION"
else
    VERSION=$(resolve_latest_version)
fi

# Tampilan header
echo "================================================="
echo "       AGY Mobile IDE Pro - One-Line Installer  "
echo "================================================="
echo "Versi target: $VERSION"
echo "Mulai mengunduh biner pra-kompilasi dari GitHub..."

# 1. Deteksi OS dan Arsitektur CPU
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"
BINARY_NAME="mobile-agy"

mobile_agy_pids() {
    # Jangan gunakan `pkill -f mobile-agy` karena dapat mematikan wrapper `agy-mobile update` sendiri.
    pgrep -x mobile-agy 2>/dev/null || true
    pgrep -x mobile-agy.exe 2>/dev/null || true
}

mobile_agy_running() {
    [ -n "$(mobile_agy_pids)" ]
}

stop_mobile_agy() {
    local pids
    pids=$(mobile_agy_pids)
    if [ -n "$pids" ]; then
        kill $pids 2>/dev/null || true
    fi
}

case "$OS" in
    linux)
        case "$ARCH" in
            x86_64|amd64)
                BINARY_URL="https://github.com/gilangji/agy-mobile/releases/download/${VERSION}/mobile-agy-linux-amd64"
                ;;
            aarch64|arm64)
                BINARY_URL="https://github.com/gilangji/agy-mobile/releases/download/${VERSION}/mobile-agy-linux-arm64"
                ;;
            *)
                echo "Error: Arsitektur CPU $ARCH tidak didukung untuk Linux."
                exit 1
                ;;
        esac
        ;;
    darwin)
        case "$ARCH" in
            x86_64|amd64)
                BINARY_URL="https://github.com/gilangji/agy-mobile/releases/download/${VERSION}/mobile-agy-darwin-amd64"
                ;;
            arm64)
                BINARY_URL="https://github.com/gilangji/agy-mobile/releases/download/${VERSION}/mobile-agy-darwin-arm64"
                ;;
            *)
                echo "Error: Arsitektur CPU $ARCH tidak didukung untuk macOS."
                exit 1
                ;;
        esac
        ;;
    mingw*|msys*|cygwin*|windows*)
        # Lingkungan Windows menggunakan Bash (Git Bash / MSYS2)
        BINARY_URL="https://github.com/gilangji/agy-mobile/releases/download/${VERSION}/mobile-agy-windows-amd64.exe"
        BINARY_NAME="mobile-agy.exe"
        ;;
    *)
        echo "Error: Sistem Operasi $OS tidak didukung."
        exit 1
        ;;
esac

# 2. Membuat folder instalasi (Memastikan tidak membuat folder bertingkat jika sudah berada di folder mobile-ide)
if [ "$(basename "$(pwd)")" != "mobile-ide" ]; then
    INSTALL_DIR="mobile-ide"
    mkdir -p "$INSTALL_DIR"
    cd "$INSTALL_DIR"
fi

generate_scripts() {
    # Membuat/memperbarui skrip start.sh
    cat <<'EOT' > start.sh
#!/bin/bash
# Tambahkan ~/.local/bin ke PATH jika belum ada
if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
    export PATH="$HOME/.local/bin:$PATH"
fi

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

if [ -f ./mobile-agy.exe ]; then
    ./mobile-agy.exe
else
    ./mobile-agy
fi
EOT
    chmod +x start.sh

    # Membuat/memperbarui skrip update.sh agar pengguna mudah menjalankan pembaruan
    cat <<'EOT' > update.sh
#!/bin/bash
set -e
TARGET_VERSION="${1:-${VERSION:-latest}}"
INSTALLER_TMP="${TMPDIR:-/tmp}/mobile-agy-install.sh"
echo "Mulai menjalankan pembaruan Mobile IDE ke versi: $TARGET_VERSION"
curl -H 'Cache-Control: no-cache' -fsSL 'https://raw.githubusercontent.com/gilangji/agy-mobile/main/install.sh' -o "$INSTALLER_TMP"
exec env VERSION="$TARGET_VERSION" bash "$INSTALLER_TMP"
EOT
    chmod +x update.sh

    # Membuat perintah global 'agy-mobile'
    CURRENT_AGY_MOBILE_PATH=$(which agy-mobile 2>/dev/null || echo "")
    if [ -n "$CURRENT_AGY_MOBILE_PATH" ] && [ -w "$CURRENT_AGY_MOBILE_PATH" ]; then
        TARGET_WRAPPER="$CURRENT_AGY_MOBILE_PATH"
    else
        TARGET_WRAPPER="$HOME/.local/bin/agy-mobile"
        mkdir -p "$HOME/.local/bin"
    fi
    echo "Membuat skrip wrapper global 'agy-mobile' di $TARGET_WRAPPER..."
    ABS_INSTALL_DIR="$(pwd)"
    cat <<EOF > "$TARGET_WRAPPER"
#!/bin/bash
# Antigravity Mobile IDE Wrapper CLI

INSTALL_DIR="$ABS_INSTALL_DIR"

mobile_agy_pids() {
    pgrep -x mobile-agy 2>/dev/null || true
    pgrep -x mobile-agy.exe 2>/dev/null || true
}

mobile_agy_running() {
    [ -n "\$(mobile_agy_pids)" ]
}

stop_mobile_agy() {
    local pids
    pids=\$(mobile_agy_pids)
    if [ -n "\$pids" ]; then
        kill -9 \$pids 2>/dev/null || true
    fi
    fuser -k 8080/tcp 2>/dev/null || true
}

case "\$1" in
    start)
        echo "Menjalankan Mobile IDE..."
        if mobile_agy_running; then
            echo "Mobile IDE sudah berjalan."
        else
            cd "\$INSTALL_DIR"
            if command -v setsid &>/dev/null; then
                setsid ./start.sh > server.log 2>&1 &
            else
                nohup ./start.sh > server.log 2>&1 &
            fi
            sleep 2
            if mobile_agy_running; then
                echo "Mobile IDE berhasil dijalankan."
            else
                echo "Gagal menjalankan Mobile IDE. Periksa \$INSTALL_DIR/server.log untuk melihat detail error."
            fi
        fi
        ;;
    stop)
        echo "Menghentikan Mobile IDE..."
        stop_mobile_agy
        echo "Mobile IDE telah dihentikan."
        ;;
    restart)
        echo "Memulai ulang Mobile IDE..."
        stop_mobile_agy
        sleep 1
        cd "\$INSTALL_DIR"
        if command -v setsid &>/dev/null; then
            setsid ./start.sh > server.log 2>&1 &
        else
            nohup ./start.sh > server.log 2>&1 &
        fi
        sleep 2
        echo "Mobile IDE telah berhasil dimuat ulang."
        ;;
    status)
        PID=\$(mobile_agy_pids)
        if [ -n "\$PID" ]; then
            echo "========================================="
            echo "        Status Mobile IDE: BERJALAN      "
            echo "========================================="
            echo "Versi    : $VERSION"
            echo "PID      : \$PID"
            if [ -f "\$INSTALL_DIR/.env" ]; then
                PORT=\$(grep -E "^PORT=" "\$INSTALL_DIR/.env" | cut -d'=' -f2)
                PASSWORD=\$(grep -E "^PASSWORD=" "\$INSTALL_DIR/.env" | cut -d'=' -f2)
                echo "Port     : \$PORT"
                echo "Kata Sandi: \$PASSWORD"
                echo "Alamat   : http://localhost:\$PORT"
            fi
            echo "========================================="
        else
            echo "========================================="
            echo "        Status Mobile IDE: BERHENTI      "
            echo "========================================="
        fi
        ;;
    logs)
        echo "=== Log Otentikasi Mobile IDE ==="
        if [ -f "\$INSTALL_DIR/server.log" ]; then
            grep -i "\[AUTH" "\$INSTALL_DIR/server.log" | tail -n 100
        else
            echo "Tidak ada berkas server.log di \$INSTALL_DIR"
        fi
        ;;
    log)
        if [ -f "\$INSTALL_DIR/server.log" ]; then
            if [ "\$2" == "-f" ] || [ "\$2" == "follow" ]; then
                tail -f "\$INSTALL_DIR/server.log"
            else
                tail -n 100 "\$INSTALL_DIR/server.log"
            fi
        else
            echo "Tidak ada berkas server.log di \$INSTALL_DIR"
        fi
        ;;
    update)
        TARGET_VERSION="\${2:-latest}"
        INSTALLER_TMP="\${TMPDIR:-/tmp}/mobile-agy-install.sh"
        echo "Memperbarui Mobile IDE ke versi \$TARGET_VERSION..."
        curl -H 'Cache-Control: no-cache' -fsSL 'https://raw.githubusercontent.com/gilangji/agy-mobile/main/install.sh' -o "\$INSTALLER_TMP"
        exec env VERSION="\$TARGET_VERSION" bash "\$INSTALLER_TMP"
        ;;
    install-version)
        if [ -z "\$2" ]; then
            echo "Penggunaan: agy-mobile install-version <tag>"
            echo "Contoh: agy-mobile install-version v1.4.1"
            exit 1
        fi
        TARGET_VERSION="\$2"
        INSTALLER_TMP="\${TMPDIR:-/tmp}/mobile-agy-install.sh"
        echo "Menginstal Mobile IDE versi \$TARGET_VERSION..."
        curl -H 'Cache-Control: no-cache' -fsSL 'https://raw.githubusercontent.com/gilangji/agy-mobile/main/install.sh' -o "\$INSTALLER_TMP"
        exec env VERSION="\$TARGET_VERSION" bash "\$INSTALLER_TMP"
        ;;
    releases)
        curl -fsSL "https://api.github.com/repos/gilangji/agy-mobile/releases?per_page=30" | grep '"tag_name":' | sed -E 's/.*"tag_name": "([^"]+)".*/\1/'
        ;;
    uninstall)
        echo "Menghentikan Mobile IDE..."
        stop_mobile_agy
        CURRENT_AGY_MOBILE_PATH=\$(which agy-mobile 2>/dev/null || echo "")
        if [ -n "\$CURRENT_AGY_MOBILE_PATH" ] && [ -w "\$CURRENT_AGY_MOBILE_PATH" ]; then
            rm -f "\$CURRENT_AGY_MOBILE_PATH"
            echo "Menghapus perintah global '\$CURRENT_AGY_MOBILE_PATH'."
        else
            rm -f "\$HOME/.local/bin/agy-mobile"
            echo "Menghapus perintah global 'agy-mobile'."
        fi

        if [ -c /dev/tty ]; then
            read -p "Apakah Anda yakin ingin menghapus direktori instalasi (\$INSTALL_DIR)? (y/N): " choice < /dev/tty
        else
            choice="n"
        fi
        if [[ "\$choice" =~ ^[Yy]$ ]]; then
            rm -rf "\$INSTALL_DIR"
            echo "Direktori instalasi (\$INSTALL_DIR) berhasil dihapus."
        else
            echo "Direktori instalasi tidak dihapus."
        fi
        echo "Mobile IDE berhasil di-uninstall."
        ;;
    *)
        echo "Penggunaan: agy-mobile {start|stop|restart|status|log|logs|update [versi]|install-version <versi>|releases|uninstall}"
        exit 1
        ;;
esac
EOF
    chmod +x "$TARGET_WRAPPER"
}

# 3. Mengunduh biner baru terlebih dahulu (Download first to minimize downtime!)
echo "Mengunduh biner untuk OS: $OS ($ARCH)..."
echo "Alamat URL: $BINARY_URL"

# Mengunduh ke file sementara (.tmp) terlebih dahulu
TEMP_BINARY="${BINARY_NAME}.tmp"
rm -f "$TEMP_BINARY"

if ! curl -fL --no-progress-meter "$BINARY_URL" -o "$TEMP_BINARY"; then
    echo "Informasi: Biner pra-kompilasi belum tersedia di Release $VERSION repositori gilangji/agy-mobile."
    echo "Mengunduh biner cadangan (fallback)..."
    FALLBACK_URL="$(echo "$BINARY_URL" | sed 's|gilangji/agy-mobile|sodikinnaa/go-agy-ide|g')"
    if ! curl -fL --no-progress-meter "$FALLBACK_URL" -o "$TEMP_BINARY"; then
        if command -v go >/dev/null 2>&1; then
            echo "Mengompilasi biner secara otomatis menggunakan Go..."
            CGO_ENABLED=0 go build -o "$TEMP_BINARY" .
        else
            echo "================================================="
            echo "ERROR: Gagal mengunduh biner dari GitHub!"
            echo "Silakan periksa koneksi internet atau buat Release di GitHub gilangji/agy-mobile."
            echo "================================================="
            exit 1
        fi
    fi
fi

if [[ "$BINARY_NAME" != *.exe ]]; then
    chmod +x "$TEMP_BINARY"
fi

# 4. Memperbarui biner menggunakan metode Hot-Swap (Mencegah 'text file busy' dan zero downtime)
echo "Melakukan hot-swap biner..."
if [[ "$BINARY_NAME" == *.exe ]]; then
    # Di Windows, menghentikan proses terlebih dahulu karena sistem mengunci file yang sedang berjalan
    taskkill //F //IM mobile-agy.exe 2>/dev/null || true
    mv -f "$TEMP_BINARY" "$BINARY_NAME" 2>/dev/null || true
else
    # Di Linux/Unix, mengganti nama file yang sedang berjalan terlebih dahulu (diizinkan oleh OS)
    if [ -f "$BINARY_NAME" ]; then
        mv -f "$BINARY_NAME" "${BINARY_NAME}.old" 2>/dev/null || true
    fi
    # Memindahkan biner baru ke lokasi utama
    mv -f "$TEMP_BINARY" "$BINARY_NAME"

    # Menghentikan proses lama yang berjalan sebagai .old tanpa mematikan wrapper `agy-mobile update`.
    stop_mobile_agy
    sleep 0.2

    # Hapus file .old (akan dihapus setelah proses lama berhenti)
    rm -f "${BINARY_NAME}.old" 2>/dev/null || true
fi

# 5. Memeriksa, Menginstal, serta Memperbarui Google Antigravity CLI (agy)
echo "Memeriksa Google Antigravity CLI (agy)..."
if ! command -v agy &> /dev/null && [ ! -f "$HOME/.local/bin/agy" ]; then
    echo "Google Antigravity CLI (agy) tidak ditemukan. Mulai mengunduh dan menginstal..."
    if ! curl -fsSL https://antigravity.google/cli/install.sh | bash; then
        echo "Peringatan: Gagal menginstal Antigravity CLI secara otomatis."
        echo "Anda dapat mencoba menginstal manual menggunakan perintah:"
        echo "  curl -fsSL https://antigravity.google/cli/install.sh | bash"
    fi
else
    echo "Google Antigravity CLI (agy) telah terinstal. Mencoba memperbarui ke versi terbaru..."
    AGY_BIN="agy"
    if [ -f "$HOME/.local/bin/agy" ]; then
        AGY_BIN="$HOME/.local/bin/agy"
    fi
    $AGY_BIN update || echo "Peringatan: Gagal memperbarui Antigravity CLI."
fi

# Tambahkan ~/.local/bin ke PATH jika belum ada dalam sesi ini
if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
    export PATH="$HOME/.local/bin:$PATH"
fi

# 6. Pengaturan workspaces.json awal
if [ ! -f "workspaces.json" ]; then
    cat <<EOT > workspaces.json
{
  "active": "$(pwd)",
  "list": [
    "$(pwd)"
  ]
}
EOT
fi

# 7. Pengaturan Konfigurasi (.env) - Mendukung Fresh Install atau Update
IS_UPDATE=false
if [ -f .env ]; then
    IS_UPDATE=true
    PORT=$(grep -E "^PORT=" .env | cut -d'=' -f2 || echo "8080")
    GEN_PASSWORD=$(grep -E "^PASSWORD=" .env | cut -d'=' -f2 || echo "AgyPass123")
    echo ""
    echo "Menemukan file konfigurasi .env (Mode Pembaruan)..."
    echo "Menggunakan port lama       : $PORT"
    echo "Menggunakan kata sandi lama : $GEN_PASSWORD"
else
    # Fresh Install: Meminta Port Keinginan Pengguna
    echo ""
    echo "-------------------------------------------------"
    if [ -c /dev/tty ]; then
        read -p "Masukkan Port untuk server Mobile IDE (Default: 8080): " USER_PORT < /dev/tty
    else
        USER_PORT=""
    fi

    PORT="8080"
    if [ -n "$USER_PORT" ]; then
        if [[ "$USER_PORT" =~ ^[0-9]+$ ]]; then
            PORT="$USER_PORT"
        else
            echo "Format port salah. Menggunakan port default 8080."
        fi
    fi

    # Generate kata sandi keamanan acak (12 karakter)
    GEN_PASSWORD=$(LC_ALL=C tr -dc A-Za-z0-9 </dev/urandom | head -c 12 2>/dev/null || echo "AgyPass123")
fi

# Dapatkan alamat DBUS dari sesi atau socket default
DBUS_ADDR="$DBUS_SESSION_BUS_ADDRESS"
if [ -z "$DBUS_ADDR" ]; then
    MY_UID=$(id -u)
    if [ -S "/run/user/$MY_UID/bus" ]; then
        DBUS_ADDR="unix:path=/run/user/$MY_UID/bus"
    fi
fi

# Tulis/perbarui file konfigurasi .env tanpa menghapus pengaturan lainnya
touch .env
set_env_var() {
    local key="$1"
    local value="$2"
    local escaped
    escaped=$(printf '%s' "$value" | sed 's/[&/\\]/\\&/g')
    if grep -qE "^${key}=" .env; then
        sed -i.bak "s/^${key}=.*/${key}=${escaped}/" .env && rm -f .env.bak
    else
        echo "${key}=${value}" >> .env
    fi
}

set_env_var "PORT" "$PORT"
set_env_var "PASSWORD" "$GEN_PASSWORD"
if [ -n "$DBUS_ADDR" ]; then
    set_env_var "DBUS_SESSION_BUS_ADDRESS" "$DBUS_ADDR"
fi

# Membuat/memperbarui seluruh skrip start.sh, update.sh, dan agy-mobile
generate_scripts

# 8. Menjalankan server di background
echo "Menjalankan server Mobile IDE di port: $PORT..."
if command -v setsid &>/dev/null; then
    setsid ./start.sh > server.log 2>&1 &
else
    nohup ./start.sh > server.log 2>&1 &
fi

# Menunggu 2 detik untuk memeriksa apakah server berhasil berjalan
sleep 2

# Periksa apakah proses masih berjalan
SERVER_RUNNING=false
if [[ "$BINARY_NAME" == *.exe ]]; then
    if pgrep -x mobile-agy.exe > /dev/null 2>&1; then
        SERVER_RUNNING=true
    fi
else
    if pgrep -x mobile-agy > /dev/null 2>&1; then
        SERVER_RUNNING=true
    fi
fi

# 9. Ringkasan Instalasi
echo "================================================="
echo "                INSTALASI SUKSES!                "
echo "================================================="
echo "Mobile IDE berhasil disetup di folder: $(pwd)"
echo "-------------------------------------------------"
echo "Port Server        : $PORT"
echo "Kata Sandi Akses   : $GEN_PASSWORD"
echo "-------------------------------------------------"

if [ "$SERVER_RUNNING" = true ]; then
    echo "Server telah berjalan di background!"
    echo "Buka browser dan akses alamat berikut:"
    echo "  http://localhost:$PORT"
    echo ""
    echo "Untuk memeriksa log server, ketik:"
    echo "  cat server.log"
else
    echo "Server gagal berjalan otomatis (kemungkinan port $PORT sudah digunakan)."
    echo "Anda dapat menjalankan server secara manual:"
    echo "  ./start.sh"
fi

echo ""
echo "Catatan: Port serta Kata Sandi Akses telah disimpan dalam file '.env'."
echo "Anda dapat mengubah file '.env' untuk kustomisasi."
echo "================================================="
