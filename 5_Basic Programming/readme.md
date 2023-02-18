# (5) Basic Programming

## Golang
Go atau Golang adalah bahasa pemrograman general-purpose yang memudahkan pembuatan, sederhana, andal, dan perangkat lunak yang efisien. adalah bahasa yang bagus untuk menulis program tingkat rendah yang menyediakan layanan ke sistem lain, yang disebut pemrograman sistem. Contoh program yang dapat dibuat dengan Go :
Program Aplikasi
    - E-Commerce
    - Pemutar musik
    - Aplikasi Media Sosial
Program Sistem
    - APIs
    - Game Engine
    - aplikasi CLI

Alasan menggunakan Go :
    1.	Menyenangkan dan Membuat Pemrograman Menyenangkan (sederhana)
    2.	Gabungkan yang terbaik dari keduanya.
        - Mengkompilasi bahasa yang diketik secara statis (seperti C) dan bahasa Dinamis (seperti JavaScript).
        - Go memiliki nuansa bahasa skrip yang lebih ringan tetapi dikompilasi.
    3.	Sintaks ringan.
    4.	Concurrency bawaan.
        - Skala besar, komputasi jaringan, seperti pencarian web Google
        - Perangkat keras multi-core
    5.	Sumber Terbuka.
    6.	Digunakan oleh Perusahaan Besar

## Variabel
Variabel digunakan untuk menyimpan informasi dalam program komputer, mereka menyediakan cara pelabelan data dengan nama deskriptif dan memiliki tipe data (boolean, numerik, string). Secara garis besar tipe data pada Golang ada 3 yaitu boolean, numerik, string.
- Boolean memiliki nilai true dan false
- Numerik dibagi menjadi 3 yaitu integer (bilangan bulat), float (bilangan desimal), complex.
- String merupakan setiap karakter yang diapit “”.
Untuk deklarasi variabel ketik var diikuti dengan nama variabel dan tipe datanya, contoh:
```var nama string```
sedangkan untuk deklarasi konstanta menggunakan awalan const, contoh:
```const nama string```


## Operator
### Arithmetic
Operator aritmatika melakukan operasi penjumlahan, pengurangan, perkalian, pembagian, eksponensial, dan modulus. Beberapa operator aritmatika:
- \+ -> Penjumlahan
- \- -> Pengurangan
- / -> Pembagian
- \* -> Perkalian
- % -> Sisa bagi
- ++ -> Increment
- -- -> Decrement

### Comparison
Operator perbandingan digunakan dalam pernyataan logis untuk menentukan kesetaraan atau perbedaan antara variabel atau nilai. Beberapa operator perbandingan:
- == -> sama dengan
- != -> tidak sama dengan
- > -> lebih dari
- < -> kurang dari
- >= -> lebih dari atau sama dengan
- <= -> kurang dari atau sama dengan

### Logical
Operator Logika digunakan untuk membandingkan dua nilai yang memiliki hasil dalam Boolean (yang bernilai true atau false). Beberapa operator logika:
- && -> dan
- || -> atau
- ! -> kebalikan

### Bitwise
Operator bitwise adalah karakter yang mewakili tindakan (operasi bitwise) yang akan dilakukan pada bit tunggal. Beberapa operator bitwise:
- & -> bitwise DAN
- | -> bitwise ATAU
- ^ -> bitwise XOR
- ~ -> bitwise tidak
- << -> geser ke kiri
- >> -> geser ke kanan

### Assignment
Operator penugasan memberikan nilai operan kanannya ke variabel, properti, atau elemen pengindeks yang diberikan oleh operan kirinya. Beberapa operator penugasan:
- =
- +=
- -=
- *=
- /=
- %= 
- dll.

### Miscellanous
Beberapa operator lain/lainnya:
- & -> Alamat
- * -> pointer ke variabel


## Control Structures
Hanya ada beberapa struktur kontrol di Go. untuk percabangan kita menggunakan if dan switch, untuk menulis loop kita menggunakan kata kunci for.
