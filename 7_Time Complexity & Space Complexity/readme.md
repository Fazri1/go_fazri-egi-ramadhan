# (7) Time Complexity & Space Complexity

## Time Complexity
Penggunaan time complexity memudahkan untuk memperkirakan waktu berjalan suatu program. Kompleksitas dapat dilihat sebagai jumlah maksimum operasi primitif yang dapat dijalankan oleh suatu program. Salah satu contoh time complexity adalah Big-O Notation.

### Macam Big-O Notation
#### Constant time - O(1)
Tingkat kompleksitas bersifat konstan berapapun jumlah data yang di proses. Ini adalah nilai unutk algoritma yang paling bagus.

#### Linear time - O(n)
Tingkat kompleksitas berbanding lurus dengan banyaknya data. Sering disebut dengan tingkat kompleksitas linear atau aritmatika.

#### Logarithmic time - O(log n)
Tingkat kompleksitas akan berbanding lurus dengan log dari banyaknya jumlah data. Apabila ada algoritma dengan komplesitas ini maka algoritma yang digunakan sangat bagus.

#### Quadratic time - O(n^2)
Sama dengan tingkat kompleksitas O(n). Jadi konstanta tidak dimasukan dalam penghitungan tingkat kompleksitas.

#### Exponential and factorial time
Perlu diketahui bahwa ada jenis lain dari kompleksitas waktu seperti waktu faktorial O(n!) dan waktu eksponensial O(2^n). Algoritma dengan kompleksitas seperti itu hanya dapat menyelesaikan masalah untuk nilai n yang sangat kecil, karena akan memakan waktu terlalu lama untuk dieksekusi untuk nilai n yang besar.

### Time Limit
Saat ini, rata-rata komputer dapat melakukan 10^8 operasi dalam waktu kurang dari satu detik. Batas waktu yang ditetapkan untuk tes online biasanya dari 1 hingga 10 detik.

- n <= 1.000.000, kompleksitas waktu yang diharapkan adalah O(n) atau O(n log n)
- n <= 10.000, kompleksitas waktu yang diharapkan adalah O(nA2)
- n <= 500, kompleksitas waktu yang diharapkan adalah O(nA3)

Tentu saja, batasan ini tidak tepat. Itu hanya perkiraan, dan akan bervariasi tergantung pada tugas tertentu.

## Space Complexity
Batas memori memberikan informasi tentang kompleksitas ruang yang diharapkan. Anda dapat memperkirakan jumlah variabel yang dapat Anda deklarasikan dalam program Anda. Singkatnya, jika Anda memiliki jumlah variabel yang konstan, Anda juga memiliki kompleksitas ruang yang konstan.