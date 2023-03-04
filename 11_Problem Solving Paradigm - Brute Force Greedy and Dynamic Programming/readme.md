# (11) Problem Solving Paradigm - Brute Force, Greedy and Dynamic Programming

Paradigma pemecahan masalah adalah pendekatan yang umum digunakan untuk memecahkan masalah: Complete Search (alias Brute Force), Divide and Conquer, pendekatan Greedy, dan Pemrograman Dinamis. Setiap masalah perlu kita selesaikan dengan pendekatan yang sesuai.

## Complete Search (Brute Force)
Complete search (juga dikenal sebagai) Bruteforce adalah metode untuk menyelesaikan masalah dengan melintasi seluruh ruang pencarian untuk mendapatkan solusi yang diperlukan. Bruteforce terjadi ketika tidak ada algoritma lain yang tersedia. Biasanya mudah ditulis karena lugas. Secara teoritis semua masalah dapat diselesaikan dengan menggunakan pendekatan Brute Force terutama ketika Anda memiliki waktu yang tidak terbatas. Time Complexity: O (n)

## Divide & Conquer
Divide & Conquer (D&C) adalah paradigma pemecahan masalah di mana suatu masalah dibuat lebih sederhana dengan 'membagi' menjadi bagian-bagian yang lebih kecil dan kemudian menaklukkan setiap bagian. Contoh: algoritma binary search dengan time complexity O (log n).

Langkah:
Divide: membagi masalah yang besar menjadi masalah yang lebih kecil.
Conquer: ketika masalah sudah cukup kecil untuk diselesaikan, langsung selesaikan. 
Combine: jika dibutuhkan maka perlu menggabungkan solusi dari masalah-masalah yang lebih kecil menjadi solusi untuk masalah yang besar.

## Greedy
Algoritma dikatakan greedy jika membuat pilihan optimal lokal pada setiap langkah dengan harapan pada akhirnya mencapai solusi optimal global. Dalam beberapa kasus, greedy bekerja - solusinya singkat dan berjalan efisien. Beberapa algoritma greedy :
- Huffman Coding
- Activity Selection
- Dijkstra algorithm - finding shortest path given (source, destination)
- dll.

## Dynamic Programming
Pemrograman Dinamis (DP) adalah teknik algoritmik untuk memecahkan masalah optimisasi dengan memecahnya menjadi submasalah yang lebih sederhana dan memanfaatkan fakta bahwa solusi optimal untuk keseluruhan masalah bergantung pada solusi optimal untuk submasalahnya.

Mari kita ambil contoh angka Fibonacci. Seperti yang kita ketahui bersama, angka Fibonacci adalah rangkaian angka yang setiap angkanya merupakan penjumlahan dari dua angka sebelumnya. Beberapa angka Fibonacci pertama adalah O, I, I, 2, 3, 5, dan 8, dan berlanjut dari sana. 

### Dynamic Programming Characteristic
1. Submasalah yang Tumpang Tindih
Submasalah adalah versi yang lebih kecil dari masalah aslinya. Masalah apa pun memiliki submasalah yang tumpang tindih jika menemukan solusinya melibatkan penyelesaian submasalah yang sama beberapa kali.
2. Properti Substruktur Optimal
Setiap masalah memiliki sifat substruktur optimal jika solusi optimal keseluruhannya dapat dibangun dari solusi optimal submasalahnya. Untuk angka Fibonacci, seperti yang kita ketahui:
`Fib(n) = Fib(n-1) + Fib(n-2)`
Ini jelas menunjukkan bahwa masalah ukuran 'n' telah direduksi menjadi submasalah ukuran 'n-1' dan 'n-2'. Oleh karena itu, angka Fibonacci memiliki properti substruktur yang optimal.

### Dynamic Programming Method
1. Top-down dengan Memoization
Dalam pendekatan ini, kami mencoba memecahkan masalah yang lebih besar dengan mencari solusi secara rekursif untuk sub-masalah yang lebih kecil. Setiap kali kami memecahkan sub-masalah, kami menyimpan hasilnya dalam cache sehingga kami tidak menyelesaikannya berulang kali jika dipanggil berkali-kali. Sebagai gantinya, kami hanya dapat mengembalikan hasil yang disimpan. Teknik menyimpan hasil subproblem yang sudah dipecahkan ini disebut Memoisasi.
2. Bottom-up dengan Tabulasi
Tabulasi adalah kebalikan dari pendekatan top-down dan menghindari rekursi. Dalam pendekatan ini, kami menyelesaikan masalah "bottom-up" (yaitu dengan menyelesaikan semua sub-masalah terkait terlebih dahulu). Ini biasanya dilakukan dengan mengisi tabel n-dimensi. Berdasarkan hasil dalam tabel, solusi untuk masalah teratas/awal kemudian dihitung.
Tabulasi adalah kebalikan dari Memoisasi, seperti dalam Memoisasi kami memecahkan masalah dan mempertahankan peta sub-masalah yang sudah diselesaikan. Dengan kata lain, dalam memoisasi, kami melakukannya dari atas ke bawah dalam arti bahwa kami menyelesaikan masalah teratas terlebih dahulu (yang biasanya berulang ke bawah untuk menyelesaikan sub-masalah).