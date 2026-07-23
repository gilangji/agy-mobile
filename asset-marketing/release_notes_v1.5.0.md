# AGY Mobile IDE Pro v1.5.0 — Integrasi Astro AI Shorts & YouTube Clipper (Auto-Crop 9:16 & Speaker Face-Tracking)

🚀 **AGY Mobile IDE Pro v1.5.0** — Official Release

AGY Mobile IDE Pro v1.5.0 hadir dengan modul besar **Astro AI Shorts & YouTube Clipper**, integrasi **Pemotong Video Vertikal 9:16 (TikTok/Reels Ready)**, **Pelacak Wajah Pembicara Otomatis (Dynamic Active Speaker Tracking)**, **Generator Subtitle Animasi Kata-per-Kata**, **Penerjemah Bahasa Otomatis**, serta **Analisis Skor Viralitas AI (Viral Score 0–100)**.

---

## 🌟 Perubahan Utama (Release Highlights)

### 🎬 1. Integrasi Modul Spesialis Astro AI Shorts & YouTube Clipper
- **Akses Bebas Bentrok**: Modul pemotong video dapat diakses via tombol **`[🎬 AI Clipper]`** di header utama IDE sebagai *full-screen overlay modal* (`#clipper-workbench-modal`) yang sangat ringan dan tidak mengganggu penyuntingan kode maupun terminal aktif.
- **4 Sub-View Utama Interaktif**:
  1. ⚡ **Ekstraksi Klip & 9:16 Crop Studio**: Memotong video panjang YouTube menjadi klip vertikal **9:16 (TikTok/Shorts/Reels)** secara presisi lengkap dengan pemutar video interaktif, stempel waktu (*Timestamp*), dan pilihan aspek rasio (9:16, 1:1, 16:9).
  2. 👤 **Speaker Face-Tracking**: Kamera potong vertikal 9:16 secara otomatis melacak dan mengunci posisi wajah tokoh yang sedang berbicara pada video podcast multi-speaker.
  3. 💬 **Subtitle & Template Text**: Generator subtitle animasi kata-per-kata bergaya *Alex Hormozi Captions* dan penerjemah otomatis ke Bahasa Indonesia baku.
  4. 📊 **Viral Score & Hook AI**: Menganalisis potensi engagement klip serta memberikan rekomendasi judul hook penarik perhatian yang siap disalin.

---

### 📱 2. Optimalisasi Header Mobile & Pengalaman Pengguna (UX)
- **Zero-Clipping Layout Header**: Tata letak tombol di header atas (`[💾 Save]`, `[🎓 Riset Academic]`, `[🛡️ Cyber Security]`, `[🎬 AI Clipper]`, `[👤 Akun]`, `[🔍 Search]`, `[🚪 Logout]`) diatur secara fleksibel dan proporsional untuk kenyamanan penuh di layar HP Android Termux.

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

Jalankan perintah ini di Termux HP atau Terminal Anda untuk langsung memperbarui ke versi **v1.5.0**:

```bash
curl -fsSL https://raw.githubusercontent.com/gilangji/agy-mobile/main/install.sh | bash
```
