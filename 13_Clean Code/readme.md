# (13) Clean Code

Clean Code adalah istilah untuk kode yang mudah dibaca, difahami dan diubah oleh programmer.

Alasan menerapkan clean code :
- Work Collaboration
- Feature Development
- Faster Development

## Karakteristik clean code :
1. Mudah dipahami -> seperti penamaan variabel dan function.
2. Mudah dieja dan dicari
3. Singkat namun mendeskripsikan konteks
4. Konsisten dalam style penamaan
5. Hindari penambahan konteks yang tidak perlu
6. Menggunakan komentar secukupnya (seperti satu komentar untuk satu blok kode)
7. Good function
8. Gunakan konvensi
9. Formatting -> contoh formatting yang baik :
    - Lebar baris code 80 - 120 karakter.
    - Satu class 300 - 500 baris.
    - Baris code yang berhubungan saling berdekatan.
    - Dekatkan fungsi dengan pemanggilnya.
    - Deklarasi variabel berdekatan dengan penggunanya.
    - Perhatikan indentasi.
    - Menggunakan ekstensi prettier atau formatter.


## Prinsip Clean Code
### KISS (Keep It So Simple)
Hindari membuat fungsi yang dibuat untuk melakukan A, sekaligus memodifikasi B, mengecek fungsi C, dst. Best practice :
- Fungsi atau class harus kecil.
- Fungsi dibuat untuk melakukan satu tugas saja.
- Jangan gunakan terlalu banyak argumen pada fungsi.
- Harus diperhatikan untuk mencapai kondisi yang seimbang, kecil dan jumlahnya minimal.

### DRY (Don't Repeat Yourself)
Duplikasi code terjadi karena sering copy paste. Untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan secara berulang-ulang.

### Refactoring
Refactoring adalah proses restrukturisasi kode yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal. Prinsip KISS dan DRY bisa dicapai dengan cara refactor. Teknik refaktoring :
- Membuat sebuah abstraksi.
- Memecah kode dengan fungsi/class.
- Perbaiki penamaan dan lokasi kode.
- Deteksi kode yang memiliki duplikasi.