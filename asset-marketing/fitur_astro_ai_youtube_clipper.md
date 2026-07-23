# Dokumentasi Fitur Astro AI Shorts & YouTube Clipper (Auto-Crop 9:16 & Subtitle Generator)

Dokumen ini berisi spesifikasi teknis dan panduan operasional modul **Astro AI Shorts & YouTube Clipper** yang terintegrasi pada **AGY Mobile IDE** dengan dukungan penuh 100% Bahasa Indonesia.

---

## 🎬 Ringkasan Fitur

**Astro AI Shorts & YouTube Clipper** adalah modul pemotong video YouTube cerdas yang memanfaatkan kecerdasan buatan (*Artificial Intelligence*) untuk memangkas video panjang YouTube menjadi klip vertikal **9:16 (TikTok, YouTube Shorts, Instagram Reels Ready)**, mendeteksi wajah pembicara secara dinamis (*Dynamic Active Speaker Tracking*), dan menghasilkan subtitle/caption animasi unik secara otomatis.

Modul ini disajikan sebagai *full-screen overlay modal* (`#clipper-workbench-modal`) yang responsif, sangat ringan, dan bebas bentrok dengan IDE utama.

---

## 🛠️ 4 Sub-View Utama

### 1. ⚡ Ekstraksi Klip & 9:16 Crop Studio
- **Input URL YouTube**: Menempelkan tautan video YouTube (`https://youtu.be/...`).
- **Deteksi Momen Viral AI (Viral Highlights Extractor)**: AI secara otomatis mengekstrak 2–5 momen terbaik dengan skor viralitas (contoh: `Viral Score: 98/100`).
- **Pemutar Video & Jump Timestamp**: Mengeklik klip secara otomatis memutar dan melompat ke detik yang sesuai di pemutar video.
- **Pilihan Aspek Rasio**: **9:16 Vertikal (TikTok/Shorts)**, **1:1 Persegi (Instagram Post)**, atau **16:9 Asli**.

### 2. 👤 Speaker Face-Tracking (Pemandu Kamera Wajah Otomatis)
- **Active Speaker Tracking**: Kamera potong 9:16 secara otomatis melacak dan memfokuskan posisi wajah tokoh yang sedang berbicara pada video podcast multi-speaker.
- **Mode Stacked 9:16 (Tutorial/Koding)**: Memotong area kodingan/terminal di bagian atas dan wajah pengajar di bagian bawah secara bertumpuk.

### 3. 💬 Subtitle & Template Text Generator
- **Auto-Captions Animated Text**: Generator subtitle animasi kata-per-kata bergaya *Alex Hormozi Captions*.
- **Penerjemah Bahasa Otomatis**: Menerjemahkan transkrip video berbahasa Inggris menjadi Bahasa Indonesia ilmiah yang rapi.

### 4. 📊 Viral Score & Hook AI Analyzer
- **Rekomendasi Judul Hook Viral**: Menghasilkan opsi judul dan teks penarik perhatian (*Hook*) yang paling berpotensi FYP di TikTok dan Reels.

---

## 🚀 Cara Penggunaan Singkat
1. Klik tombol **`[🎬 AI Clipper]`** di header atas IDE.
2. Tempel URL YouTube Anda pada kolom yang tersedia.
3. Klik **`[⚡ Ekstrak Klip Shorts AI]`**.
4. Pilih klip momen terbaik, lalu klik **`[Unduh Klip 9:16]`**.
