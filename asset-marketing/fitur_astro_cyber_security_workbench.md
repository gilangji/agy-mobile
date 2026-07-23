# Dokumentasi Fitur Astro Cyber Security Workbench (Autonomous Security Inspector)

Dokumen ini berisi spesifikasi teknis dan panduan fitur **Astro Cyber Security Workbench (Pusat Audit Keamanan Defensif & Analisis Sistem)** yang terintegrasi pada **AGY Mobile IDE** dengan dukungan penuh 100% Bahasa Indonesia.

---

## 🛡️ Ringkasan Fitur

**Astro Cyber Security Workbench** adalah modul audit keamanan defensif terpadu yang dirancang khusus untuk mempermudah pemeriksaan **Keamanan Kode (SAST), Kebocoran Rahasia (Secrets), Keamanan Web Headers, Validasi SSL/TLS, Audit Dependensi (CVE), serta Pembuatan Laporan Mitigasi Keamanan**.

Modul ini disajikan sebagai *full-screen overlay modal* (`#security-workbench-modal`) yang responsif, sangat ringan (konsumsi RAM ~25MB), dan bebas bentrok dengan lingkungan penyuntingan kode maupun terminal utama pada AGY Mobile IDE.

---

## 🛠️ 5 Sub-View Utama

### 1. ⚡ Auto-Audit Pro (Autonomous Security Inspector)
- **Eksekusi Satu-Klik (`[🚀 Jalankan Auto-Audit Keamanan]`)**: Menjalankan seluruh tahapan pemeriksaan keamanan secara otomatis dari A sampai Z.
- **Indikator Progres Realtime**: Menampilkan persentase dan status pemindaian tahap demi tahap.
- **Skor Keamanan Proyek (Security Health Scorecard)**: Menampilkan skor kesehatan keamanan (contoh: `96/100 - Sangat Aman`) beserta jumlah rincian risiko (*Critical, High, Medium, Low*).
- **Tabel Temuan Audit**: Menampilkan hasil pemeriksaan secara interaktif beserta tombol opsi perbaikan instan.

### 2. 🔑 Code & Secret Scanner (Static Code Analysis - SAST)
- **Secret Leak Detector**: Memindai seluruh berkas workspace untuk mendeteksi *hardcoded secrets* (seperti API Keys, AWS Credentials, Private Keys, Password Database, JWT Tokens).
- **Risk Pattern Analysis**: Analisis statis sintaksis kode untuk mencegah pola kodingan rentan (*SQL Injection*, *Insecure Deserialization*, *Path Traversal*).

### 3. 🌐 Web & Security Headers Inspector
- **Audit Target URL**: Memeriksa kelengkapan *Security Headers* pada aplikasi web target (`http://localhost:8080` atau URL publik).
- **Evaluasi 4 Header Utama**:
  - *Content-Security-Policy (CSP)*: Pencegahan serangan XSS & injeksi skrip.
  - *Strict-Transport-Security (HSTS)*: Enkripsi HTTPS wajib.
  - *X-Frame-Options*: Perlindungan dari serangan Clickjacking.
  - *X-Content-Type-Options*: Blocking MIME Sniffing.

### 4. 📦 Dependensi & CVE Scanner
- **Package Manifest Audit**: Memeriksa file manifest dependensi proyek (`package.json`, `go.mod`, `requirements.txt`).
- **CVE Database Lookup**: Memastikan pustaka yang digunakan bebas dari catatan kerentanan publik (CVE ID).

### 5. 📊 Laporan Audit Keamanan & Mitigasi
- **Ringkasan Eksekutif**: Menampilkan rangkuman resmi status keamanan sistem.
- **Panduan Pengetatan (Hardening Guide)**: Panduan rekayasa pengetatan konfigurasi keamanan.
- **Ekspor Dokumen Satu-Klik**: Mengunduh laporan hasil audit ke dalam berkas **Markdown (`.md`)** atau **PDF**.

---

## 🎨 Aksesibilitas & Responsivitas Mobile HP
- **Bebas Bentrok**: Dapat dibuka via tombol `[🛡️ Cyber Security]` di header utama kapan saja tanpa mengganggu editor kode atau proses terminal yang sedang berjalan.
- **Tampilan Dark Mode Cyan/Teal Cyber Shield**: Visual modern dengan kombinasi warna `#070913`, aksen cyan, teal, dan font mono yang nyaman dibaca di layar HP Android Termux maupun Desktop.
