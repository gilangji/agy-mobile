# AGY Mobile IDE Pro v1.3.0 — Integrasi Astro Research Workbench & Visual Graph Pengetahuan

🚀 **AGY Mobile IDE Pro v1.3.0** — Official Release

AGY Mobile IDE Pro v1.3.0 hadir dengan penambahan modul besar **Astro Research Workbench (Asisten Pakar Riset Prof. AI)**, integrasi **Open Knowledge Maps & Visual Graph Pengetahuan**, dukungan **Unggah Dokumen Terlampir (PDF, Word, TXT, MD)** pada kolom chat, serta fitur **Rendering Formula LaTeX Matematika (KaTeX)** & ekspor PDF instan.

---

## 🌟 Perubahan Utama (Release Highlights)

### 🎓 1. Integrasi Modul Spesialis Astro Research Workbench (Asisten Riset Academic)
- **Akses Bebas Bentrok**: Modul riset dapat diakses via tombol **`[Riset Academic]`** di header utama IDE sebagai *full-screen overlay modal* (`#research-workbench-modal`) yang tidak mengganggu sesi penyuntingan kode, file tree, maupun terminal aktif.
- **5 Sub-View Utama Interaktif**:
  1. 🟢 **Dashboard Riset**: Banner Selamat Datang, **8 Alur Tahapan Bimbingan Riset Prof. AI** (Eksplorasi Topik, Kajian Pustaka, Research Gap, Metodologi & Validitas, Pengolahan & Pembahasan, Metode AK.SA.RA, Review & Publikasi, Simulasi Sidang), **8 Alat Spesialis Akademik** (SPSS/PLS, Paraphrasing, Bahasa Baku, Sintesis Pustaka, dll.), dan **Progres Struktur Naskah Skripsi/Tesis** (Bab 1–5).
  2. 📄 **Draf Bab & Editor Naskah**: Split-pane Editor Markdown + Live Preview Ilmiah dengan rendering formula LaTeX ($\sum_{i=1}^n X_i = \mu$), simpan draf otomatis ke `localStorage`, serta ekspor naskah ke format Microsoft Word (`.docx`), PDF (`.pdf`), dan LaTeX (`.tex`).
  3. 📚 **Kelola Sitasi & Referensi**: Impor metadata otomatis via DOI / BibTeX dan Generator Sitasi Multigaya (APA 7th, IEEE, Harvard, MLA 9th) dengan tombol 1-click copy.
  4. 🕸️ **Peta Pengetahuan (Open Knowledge Maps & Visual Graph)**: Pemetaan visual hubungan antar teori berbasis *Central Topic Hub*, preset topik instan (`AI in Education`, `Metodologi Riset`, `Systematic Literature Review`), serta 4 Klaster Literatur (Teori Utama, Metodologi, Aplikasi Empiris, & Research Gap).
  5. 🤖 **AI Asisten Riset (Prof. AI)**: Interface percakapan AI terdedikasi untuk bimbingan naskah akademik.

---

### 📎 2. Dukungan Unggah & Analisis Dokumen Terlampir (Paperclip Attachment)
- **Tombol Lampiran Paperclip (`[📎]`)**: Memungkinkan pengguna melampirkan berkas langsung pada kolom percakapan Prof. AI.
- **Dukungan Format Luas**: Berkas **PDF (`.pdf`)**, **Microsoft Word (`.doc`, `.docx`)**, **Teks (`.txt`)**, **Markdown (`.md`)**, **JSON (`.json`)**, dan **CSV (`.csv`)**.
- **Pratinjau Badge Visual & Analisis AI**: Menampilkan badge pratinjau berkas terlampir secara visual dan secara otomatis mengekstrak isi dokumen untuk dianalisis, diringkas, atau diparafrasa oleh Prof. AI.

---

### 📐 3. Rendering Formula LaTeX Matematika (KaTeX) & Ekspor PDF Instan
- **Integrasi KaTeX Engine**: Formula matematika LaTeX (`$$\sum_{i=1}^{n} X_i = \mu$$`) dirender secara presisi dan indah pada jendela *Live Preview Ilmiah*.
- **Client-Side PDF Generator**: Ekspor naskah ilmiah langsung menjadi berkas `.pdf` siap cetak menggunakan `html2pdf.js`.

---

### 🌐 4. PWA Offline Caching & Penyempurnaan Performa
- **PWA Service Worker (`sw.js`)**: Pembaruan cache `agy-ide-v1.3.0-pro` yang menyimpan pustaka KaTeX dan rendering ilmiah agar modul riset tetap berjalan 100% offline di lingkungan Termux Android.
- **Dukungan 100% Bahasa Indonesia Baku**: Pembaruan total seluruh antarmuka, notifikasi, dan dokumen spesifikasi pada direktori `asset-marketing/`.

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

Jalankan perintah ini di Termux HP atau Terminal Anda untuk langsung memperbarui ke versi **v1.3.0**:

```bash
curl -fsSL https://raw.githubusercontent.com/gilangji/agy-mobile/main/install.sh | bash
```
