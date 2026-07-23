# AGY Mobile IDE Pro v1.4.0 — Integrasi Astro Cyber Security Workbench & Autonomous Security Inspector

🚀 **AGY Mobile IDE Pro v1.4.0** — Official Release

AGY Mobile IDE Pro v1.4.0 hadir dengan modul besar **Astro Cyber Security Workbench (Pusat Audit Keamanan Defensif)**, integrasi **Auto-Audit Pro (Autonomous Security Inspector 1-Klik)**, **SAST & Secret Leak Scanner**, **Web Security Headers Inspector**, **Audit Dependensi (CVE)**, serta generator **Laporan Audit Keamanan (.md / .pdf)**.

---

## 🌟 Perubahan Utama (Release Highlights)

### 🛡️ 1. Integrasi Modul Spesialis Astro Cyber Security Workbench
- **Akses Bebas Bentrok**: Modul keamanan defensif dapat diakses via tombol **`[🛡️ Cyber Security]`** di header utama IDE sebagai *full-screen overlay modal* (`#security-workbench-modal`) yang sangat ringan (konsumsi RAM ~25MB) dan tidak mengganggu penyuntingan kode maupun terminal aktif.
- **5 Sub-View Utama Interaktif**:
  1. ⚡ **Auto-Audit Pro (Autonomous Inspector)**: Fitur inspeksi otomatis 1-klik (`[🚀 Jalankan Auto-Audit Keamanan]`) yang memindai seluruh berkas workspace, konfigurasi jaringan, header web, dan dependensi secara otomatis dari A sampai Z. Dilengkapi skor kesehatan keamanan proyek (Security Health Scorecard 0–100).
  2. 🔑 **Code & Secret Scanner (SAST)**: Pemindaian statis sintaksis kode untuk mendeteksi *hardcoded secrets* (seperti API Keys, Private Keys, Password, JWT Tokens) dan mencegah pola kodingan rentan (*SQL Injection*, *XSS*, *Path Traversal*).
  3. 🌐 **Web & Security Headers Inspector**: Pengujian kelengkapan *Security Headers* (`CSP`, `HSTS`, `X-Frame-Options`, `CORS Policy`) dan validasi enkripsi SSL/TLS pada peramban/server target.
  4. 📦 **Dependensi & CVE Scanner**: Memeriksa berkas manifest dependensi (`package.json`, `go.mod`, `requirements.txt`) untuk memastikan tidak ada pustaka yang rentan terhadap catatan CVE publik.
  5. 📊 **Laporan Audit Keamanan**: Dokumentasi ringkasan audit resmi lengkap dengan panduan rekayasa pengetatan (*Hardening Guide*) yang dapat diunduh instan dalam format `.md` atau `.pdf`.

---

### 🎓 2. Penyempurnaan Modul Astro Research Workbench & Prof. AI
- **Pemuatan Model AI Dinamis (Dynamic Model Detection)**: Pemilih model AI di Prof. AI secara otomatis mendeteksi dan menggunakan model agent aktif dari backend Antigravity Engine (`/api/models`).
- **Penyimpanan Riwayat Chat LocalServer**: Percakapan dengan Prof. AI tersimpan otomatis di `localStorage` memori server lokal, dilengkapi tombol 1-klik **`[🗑️ Hapus Riwayat]`**.

---

### 📱 3. Optimalisasi Tampilan Header & Aksesibilitas Mobile HP
- **100% Zero-Clipping Layout**: Penyesuaian tata letak tombol header HP Android di mode portrait/landscape sehingga tombol **`[👤 Akun]`**, **`[🛡️ Cyber Security]`**, dan **`[🎓 Riset Academic]`** tampil utuh, presisi, dan sangat nyaman diklik.

---

## 📦 Berkas Biner Rilis (Artifacts)

| Nama Berkas | Platform / Arsitektur |
|---|---|
| `mobile-agy-linux-arm64` | Android Termux (ARM64) / Linux ARM64 |
| `mobile-agy-linux-amd64` | Linux x86_64 / Server Cloud |
| `mobile-agy-darwin-arm64` | macOS Apple Silicon (M1/M2/M3/M4) |
| `mobile-agy-darwin-amd64` | macOS Intel |
| `mobile-agy-windows-amd64.exe` | Windows 64-bit |

---

## ⚡ Perintah Perbarui / Instal Cepat (One-Line Installer)

Jalankan perintah ini di Termux HP atau Terminal Anda untuk langsung memperbarui ke versi **v1.4.0**:

```bash
curl -fsSL https://raw.githubusercontent.com/gilangji/agy-mobile/main/install.sh | bash
```
