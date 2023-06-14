# (23) Unit Testing

Apa Itu Pengujian Perangkat Lunak?
Ui = (end to end) menguji interaksi antara keseluruhan melalui user interface
Integsuatu proses menganalisis item perangkat lunak untuk mendeteksi perbedaan antara kondisi yang ada dan yang dibutuhkan (mis., cacat) dan untuk mengevaluasi fitur item perangkat lunak.

Tujuan Pengujian
- mencegah regresi
- tingkat keyakinan dalam refactoring
- tingkatkan desain kode
- dokumentasi kode
- menskalakan timrasi = menguji modul atau subsistem tertentu melalui api
Unit = menguji bagian terkecil yang dapat diuji (operasi logis tunggal) dari suatu aplikasi melalui metode

Tingkat Pengujian
Ui = (end to end) menguji interaksi antara keseluruhan melalui user interface
Integrasi = menguji modul atau subsistem tertentu melalui api
Unit = menguji bagian terkecil yang dapat diuji (operasi logis tunggal) dari suatu aplikasi melalui metode

## Framework
Berdasarkan bahasa pemrograman anda, anda harus memilih framework pengujian unit. Framework menyediakan alat, dan struktur yang diperlukan untuk menjalankan pengujian secara efisien.

### Structure
2 pola biasa:
1. Pusatkan file pengujian anda di dalam folder tes.
2. Simpan file pengujian bersama dengan file produksi.

File uji (kumpulan suite uji)
Suite uji (kumpulan kasus uji)
Perlengkapan tes (pengaturan & pembongkaran)
Uji kasus (masukan, proses, keluaran)

### Runner
- alat yang menjalankan file uji
- gunakan mode jam tangan (jika ada perubahan dalam basis kode, secara otomatis jalankan tes lagi, buat tdd lebih efisien)
- pilih pelari yang bisa berlari paling cepat

### Mocking
Kasing uji anda harus mandiri, anda tidak harus menguji:
- modul pihak ke-3
- basis data
- api pihak ke-3 atau sistem eksternal lainnya
- kelas objek yang telah anda uji di tempat lain

Anda perlu membuat mock object (n) objek palsu yang mensimulasikan perilaku objek nyata

### Coverage
Pembuat kode perlu memastikan apakah mereka telah membuat cukup pengujian.
Alat cakupan menganalisis kode aplikasi saat pengujian sedang berjalan.