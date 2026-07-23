# Dokumentasi Fitur Astro Academy (Autonomous Adaptive Learning Engine)

Dokumen ini berisi spesifikasi teknis dan panduan operasional modul **Astro Academy** yang terintegrasi pada **AGY Mobile IDE** dengan dukungan penuh 100% Bahasa Indonesia.

---

## 🎓 Ringkasan Fitur

**Astro Academy** adalah pusat pembelajaran mandiri terstruktur berbasis kecerdasan buatan (*Autonomous Adaptive Learning Engine*) yang mampu memetakan kurikulum belajar berstandar internasional untuk disiplin ilmu apa saja (Psikologi, Cyber Security, Computer Science, Machine Learning, Hukum, Ekonomi, dll.) dari tingkat **Pemula (SMA/Dasar)** hingga **Doktoral / Profesor (AHLI/Riset)**.

Modul ini disajikan sebagai *full-screen overlay modal* (`#academy-workbench-modal`) yang sangat ringan (konsumsi RAM ~25MB), responsif, dan bebas bentrok dengan lingkungan IDE utama.

---

## 🛠️ 4 Sub-View Utama

### 1. ⚡ Generator Kurikulum Adaptif AI (`academy-tab-setup`)
- **Pencarian / Input Topik Kustom**: Pengguna dapat memilih preset topik atau mengetik topik apa saja yang ingin dipelajari secara mendalam (misal: *"Psikologi Kognitif & Sejarah Wilhelm Wundt"*).
- **Profil Jenjang Kedalaman**:
  - 🐣 *Dasar / Pemula (SMA / Beginner)*: Etimologi, akar sejarah, konsep dasar & analogi intuitif.
  - 🚀 *Menengah (S1 / Intermediate)*: Studi kasus praktis, dinamika cabang-cabang utama & metodologi.
  - 🎓 *Doktoral / Profesor (Advanced)*: Analisis filsafat ilmu, perbedaan paradigma tokoh & riset jurnal.
- **Generator Kurikulum 1-Klik**: Menghasilkan struktur peta jalan belajar berbab-bab secara otomatis.

### 2. 🗺️ Peta Kurikulum & Daftar Bab (`academy-tab-curriculum`)
- **Penataan Bab Terstruktur**: Setiap bab mencakup sejarah awal, etimologi (asal kata), tokoh pelopor, corak berpikir (*schools of thought*), hingga latihan praktis.
- **Navigasi Satu-Klik**: Membuka materi bab langsung di Jendela Ruang Belajar.

### 3. 📖 Jendela Ruang Belajar & Materi Interaktif (`academy-tab-learn`)
- **Sajian Materi Ilmiah Lengkap**: Pemaparan mendalam mencakup etimologi (*Psyche & Logos*), garis waktu sejarah (Yunani/Romawi -> Laboratorium Leipzig 1879 oleh Wilhelm Wundt), tokoh pelopor (Sigmund Freud, B.F. Skinner, Carl Rogers), serta mazhab utama (Psikoanalisis, Behaviorisme, Humanistik).

### 4. 📝 Uji Pemahaman & Kuis Adaptif (`academy-tab-quiz`)
- **Self-Assessment Quiz**: Uji pemahaman mandiri berbentuk kuis interaktif pilihan ganda dan studi kasus lengkap dengan ulasan jawaban otomatis.

---

## 🚀 Cara Penggunaan Singkat
1. Klik tombol **`[🎓 Astro Academy]`** di header atas IDE.
2. Ketik topik yang ingin Anda pelajari (misal: *Psikologi Umum*).
3. Pilih tingkat jenjang Anda (Dasar / Menengah / Doktoral).
4. Klik **`[Generasi Kurikulum Belajar Terstruktur]`** dan mulai membaca materi serta mengikuti kuis!
