# Dokumentasi Fitur Kontrol Panel - AGY Mobile IDE

Dokumen ini berisi spesifikasi teknis dan petunjuk operasional Kontrol Panel pada **AGY Mobile IDE** dengan dukungan penuh 100% Bahasa Indonesia.

## 📱 Ringkasan Kontrol Panel

AGY Mobile IDE menyediakan Panel Kontrol terintegrasi yang memudahkan pengembang mengelola lingkungan koding langsung dari perangkat mobile Android/Termux maupun browser.

---

## 🛠️ Fitur Utama Kontrol Panel

### 1. Manajemen Akun & Pool Otentikasi (Google OAuth & OpenAI)
- **Pool Akun Google OAuth**: Pengguna dapat menambahkan, mengalihkan (switch), dan menghapus akun Google yang terhubung untuk kuota Antigravity AI.
- **Penyedia OpenAI-Compatible**: Mengonfigurasi API Key, Endpoint kustom (seperti Ollama, vLLM, LMStudio, atau OpenAI), serta pemilihan model AI pendukung.
- **Manajemen Kata Sandi Keamanan**: Mengubah dan memperbarui kata sandi keamanan server secara instan melalui modal setelan.

### 2. Manajemen Workspace & Berkas
- **Pencari Berkas Global (Global File Search)**: Pencarian cepat berbasis nama file dan isi teks (Ctrl+P / Cmd+P) dengan pencahayaan teks pencarian.
- **Pohon Berkas Touch-Friendly**: Membuat file baru, folder baru, mengganti nama, serta menghapus dengan konfirmasi Bahasa Indonesia.
- **Bilah Alat Simbol Cepat (Touch Shortcut Toolbar)**: Menyediakan tombol cepat untuk karakter koding penting (`=`, `;`, `{`, `}`, `(`, `)`, `>`, `<`, `_`, `-`, `$`, `/`, `\`, `|`) untuk mempercepat pengetikan pada keyboard layar HP.

### 3. Pemantau Kuota & Performa
- **Ringkasan Kuota Real-time**: Menampilkan sisa kuota model (Gemini 2.5 Flash, Pro, Flash Thinking, Claude, dll.) secara visual dengan indikator persentase.
- **Penyuntingan & Terminal Interaktif**: Integrasi Terminal `xterm.js` dengan runner otomatis untuk perintah Go, Node.js, Python, dan Shell.

### 4. Sistem Pembaruan Mandiri (Self-Update System)
- **Pemeriksaan Rilis GitHub**: Memeriksa tag rilis terbaru dari repositori GitHub dan memungkinkan pembaruan/downgrade versi server dengan satu kali klik.

### 5. Modul Spesialis Astro Research Workbench (Asisten Riset Academic)
- **Overlay Layar Penuh Bebas Bentrok**: Modul khusus riset ilmiah yang dapat diakses langsung via tombol header `[Riset Academic]`.
- **5 Sub-View Utama**: Dashboard Riset, Draf Bab & Editor Naskah (dengan rendering formula LaTeX KaTeX live preview), Kelola Sitasi & Referensi (Generator APA/IEEE & Auto DOI), Peta Pengetahuan (Open Knowledge Map), dan AI Asisten Riset (Prof. AI).
- **8 Alur Tahapan & 8 Alat Spesialis**: Alur bimbingan skripsi/tesis terpandu dan perangkat analisis statistik/parafrasa Turnitin.

### 6. Modul Spesialis Astro Cyber Security Workbench (Autonomous Security Inspector)
- **Overlay Keamanan Layar Penuh**: Modul khusus audit keamanan defensif yang dapat diakses langsung via tombol header `[Cyber Security]`.
- **5 Sub-View Utama Audit Keamanan**:
  1. *Auto-Audit Pro*: Inspeksi otomatis 1-klik untuk pemindaian kode, jaringan, header web, dan dependensi secara menyeluruh.
  2. *Code & Secret Scanner (SAST)*: Memindai seluruh berkas workspace dari risiko kebocoran API Key, Private Key, JWT Token, dan Password.
  3. *Web & Security Headers Inspector*: Evaluasi kelengkapan Security Headers (CSP, HSTS, X-Frame-Options, CORS) dan SSL/TLS.
  4. *Dependensi & CVE Scanner*: Memeriksa berkas manifest (`package.json`, `go.mod`, `requirements.txt`) dari risiko pustaka rentan.
  5. *Laporan Audit Keamanan*: Pembuatan laporan audit resmi (.md / .pdf) lengkap dengan panduan rekayasa pengetatan (Hardening Guide).

---

## 📋 Pengaturan Konfigurasi Bahasa

Seluruh elemen antarmuka Kontrol Panel telah disesuaikan ke Bahasa Indonesia baku:
- **Tampilan Autentikasi (`login.html` & `login-pwd.html`)**: Halaman masuk yang bersih, aman, dan informatif.
- **Tampilan Utama IDE (`index.html`)**: Menampilkan tooltip, instruksi chat AI, dialog konfirmasi, dan notifikasi dalam Bahasa Indonesia.
- **Perlindungan Privasi**: Mengisolasi kredensial lokal dan menyembunyikan API key sensitif.

---

## 🚀 Panduan Eksekusi Perintah Terminal CLI

| Perintah | Fungsi / Kegunaan |
|---|---|
| `agy-mobile status` | Memeriksa status server dan otentikasi aktif |
| `agy-mobile logs` | Membaca log aktivitas server dan kesalahan |
| `agy-mobile update` | Memperbarui biner server ke versi rilis terbaru |
| `/credits` | Memeriksa penggunaan kredit dan kuota organisasi di CLI |
