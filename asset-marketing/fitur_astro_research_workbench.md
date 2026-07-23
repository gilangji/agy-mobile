# Dokumentasi Fitur Astro Research Workbench (Asisten Riset Academic)

Dokumen ini berisi spesifikasi teknis dan panduan fitur **Astro Research Workbench (Asisten Riset Academic / Prof. AI)** yang terintegrasi pada **AGY Mobile IDE** dengan dukungan penuh 100% Bahasa Indonesia.

---

## 🎓 Ringkasan Fitur

**Astro Research Workbench** adalah modul pakar riset terpadu yang dirancang khusus untuk mempermudah penyusunan **Skripsi, Tesis, Disertasi, dan Karya Tulis Ilmiah**. Modul ini disajikan sebagai *overlay modal* yang responsif dan tidak bertabrakan dengan lingkungan penyuntingan kode maupun terminal utama pada AGY Mobile IDE.

---

## 🛠️ 5 Sub-View Utama

### 1. 🟢 Dashboard Riset
- **Hero Welcome Card**: Sambutan interaktif dan pengenalan alur riset.
- **8 Alur Tahapan Bimbingan Riset Prof. AI**:
  1. *Eksplorasi Topik*: Analisis akar masalah utama (Why Analysis 5x).
  2. *Kajian Pustaka*: Pencarian referensi terindeks (Scholar, Scopus, SINTA).
  3. *Research Gap*: Identifikasi kesenjangan teori (PICOS) & SMART Goals.
  4. *Metodologi & Validitas*: Desain Kuantitatif, Kualitatif, Mixed-Methods & uji instrumen.
  5. *Pengolahan & Pembahasan*: Pembahasan data self-explanatory (Why & How).
  6. *Metode AK.SA.RA*: Parafrasa anti-plagiasi (AKui, parafrASa, integRASI) untuk Turnitin.
  7. *Review & Publikasi*: Simulasi Peer-Review, Think-Check-Submit & OJS Metadata.
  8. *Simulasi Sidang*: Persiapan slide PPT & simulasi tanya-jawab dosen penguji.

- **8 Alat & Fitur Spesialis Akademik**:
  1. *Paraphrasing & Anti-Plagiasi*: Ubah struktur kalimat agar unik & lolos Turnitin.
  2. *Bahasa Akademik Baku*: Penyesuaian kaidah PUEBI / EYD baku.
  3. *Format Sitasi & Referensi*: Pengaturan otomatis body citation & daftar pustaka.
  4. *Sintesis Tinjauan Pustaka*: Penggabungan ringkasan artikel komparatif.
  5. *Ekstraksi Poin Jurnal*: Ekstraksi latar belakang, metodologi, dan sampel jurnal.
  6. *Rumusan & Hipotesis*: Perumusan masalah dan hipotesis penelitian.
  7. *Analisis Output SPSS / PLS*: Penerjemahan tabel statistik ($R^2$, F-Hitung, p-value).
  8. *Matriks Tinjauan Pustaka*: Pembuatan tabel matriks sintesis komparatif otomatis.

- **Progres Struktur Naskah Skripsi / Tesis**:
  - Bab 1: Pendahuluan
  - Bab 2: Tinjauan Pustaka
  - Bab 3: Metodologi
  - Bab 4: Hasil & Pembahasan
  - Bab 5: Penutup & Saran

### 2. 📄 Draf Bab & Editor Naskah
- **Struktur Dokumen Sidebar**: Navigasi antar bab naskah (`Bab 1` s/d `Bab 5` & `Daftar Pustaka`).
- **Penyuntingan Code Split-Pane**: Editor Markdown + Live Preview Ilmiah.
- **Dukungan Formula LaTeX**: Rendering otomatis rumus matematika ($\sum_{i=1}^n X_i = \mu$) menggunakan KaTeX CDN secara real-time.
- **Pilihan Ekspor**: Ekspor naskah ke format Microsoft Word (`.docx`), PDF (`.pdf`), dan LaTeX (`.tex`).

### 3. 📚 Kelola Sitasi & Referensi
- **Impor Metadata Otomatis via DOI / BibTeX**: Cukup tempel nomor DOI (contoh: `10.1016/j.pas.2025.01`) atau teks BibTeX untuk mengisi otomatis metadata riset.
- **Generator Sitasi Multigaya**: Mendukung gaya penulisan standar **APA 7th Edition**, **IEEE**, **Harvard**, dan **MLA 9th Edition**.
- **Salin Instan**: Tombol satu klik untuk menyalin hasil sitasi ke clipboard.

### 4. 🕸️ Peta Pengetahuan (Open Knowledge Maps & Visual Graph)
- **Visual Graph Pengetahuan Interaktif**: Visualisasi peta keterkaitan antar konsep, kelompok teori, referensi kunci, dan *Research Gap* berbasis *Central Topic Hub*.
- **Tombol Preset Topik Instan**: Menyediakan preset pencarian cepat seperti `🧁 AI in Education`, `📊 Metodologi Riset`, dan `📚 Systematic Literature Review`.
- **4 Klaster Pemetaan Riset**:
  1. *Klaster 1 (Teori Utama & Fondasi)*: Kajian teori dasar dan fondasi konsep utama.
  2. *Klaster 2 (Metodologi & Pengukuran)*: Pendekatan kuantitatif/kualitatif dan instrumen uji.
  3. *Klaster 3 (Aplikasi Empiris & Studi Kasus)*: Penerapan praktis pada berbagai sampel dan konteks.
  4. *Klaster 4 (Research Gap & Peluang Kebaruan)*: Pemetaan ruang kosong literatur dan peluang kebaruan (*novelty*) riset.
- **Pencarian Dynamic Knowledge Map**: Pengguna dapat memasukkan kata kunci bidang riset kustom apa saja untuk membuat *Knowledge Graph* terstruktur secara real-time.

### 5. 🤖 AI Asisten Riset (Prof. AI) & Integrasi Antigravity Engine
- **Asisten Pakar Riset Terintegrasi Realtime**: Terhubung 100% secara langsung dengan backend **Antigravity AI Engine (`/api/chat`)** untuk menghasilkan respon ilmiah yang nyata, dinamis, dan streaming token-by-token.
- **Pemilih Model AI (Model Selector Dropdown)**: Pengguna dapat memilih model AI yang ingin digunakan secara langsung di dalam header Prof. AI (seperti *Gemini 3.5 Flash*, *Gemini 3.5 Pro*, *Claude 3.5 Sonnet*, *GPT-4o*).
- **Penyimpanan & Penghapusan Riwayat Chat LocalServer**:
  - Percakapan secara otomatis tersimpan di `localStorage` memori server lokal, sehingga riwayat chat tidak hilang saat halaman ditutup/di-refresh.
  - **Tombol `[🗑️ Hapus Riwayat]`**: Menyediakan opsi satu-klik untuk membersihkan seluruh riwayat chat lokal kapan saja.
- **Fitur Unggah Dokumen & Lampiran (Paperclip Attachment)**:
  - Mendukung pengunggahan berkas berformat **PDF (`.pdf`)**, **Microsoft Word (`.doc`, `.docx`)**, **Teks Polos (`.txt`)**, **Markdown (`.md`)**, **JSON (`.json`)**, dan **CSV (`.csv`)**.
  - **Pratinjau Badge Terlampir**: Menampilkan nama berkas dan ukuran file terlampir secara visual sebelum dikirim.
  - **Analisis Berkas Otomatis**: Berkas yang diunggah secara otomatis dibaca dan dilampirkan sebagai konteks prompt kepada Prof. AI untuk dianalisis, diringkas, atau diparafrasa.

---

## 🎨 Aksesibilitas & Responsivitas
- **Bebas Bentrok**: Dapat dibuka via tombol `[Riset Academic]` di header utama kapan saja tanpa mengganggu editor kode atau proses terminal yang sedang berjalan.
- **Tampilan Dark Mode Glassmorphism**: Visual modern dengan kombinasi warna `#090a15`, aksen emerald, cyan, dan font yang nyaman dibaca di layar HP maupun desktop.
