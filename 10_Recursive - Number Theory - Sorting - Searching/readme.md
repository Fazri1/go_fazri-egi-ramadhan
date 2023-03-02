# (10) Recursive - Number Theory - Sorting - Searching

## Recursive
Apa itu Rekursif
Rekursi adalah keadaan di mana suatu fungsi memecahkan masalah dengan memanggil sendiri berulang kali. Jika masalahnya cukup kecil, fungsi rekursi dapat menghasilkan jawaban segera. Jika masalahnya terlalu besar, fungsi akan memanggil dirinya sendiri dengan yang ruang lingkup masalah lebih kecil.

Mengapa kita butuh recursion?
- Banyak masalah lebih mudah untuk memecahkan dan mempersingkat kode saat menggunakan pendekatan rekursif. 
- Pada dasarnya, baik iteratif (misalnya dengan for loop) dan strategi rekursif melakukan sesuatu yang berulang. 
- Namun, terkadang solusi iteratif untuk suatu masalah sangat sulit dipikirkan dan memerlukan teknik khusus. 
- Dengan solusi rekursif, akan lebih mudah untuk melihat dan mendesain jalan untuk penyelesaian.

Recursion Strategy
Ada dua hal yang perlu dipikirkan saat menggunakan strategi rekursif: 
- Kasus dasar 
Apa kasus paling sederhana dari masalah ini? 
- Hubungan perulangan 
Apa hubungan rekursif dari masalah ini dengan masalah serupa yang lebih kecil?

## Number Theory
Teori bilangan adalah cabang matematika yang mempelajari bilangan bulat. Ada banyak topik dalam bidang teori bilangan yaitu Bilangan Prima, Pembagi Persekutuan Terbesar, Kelipatan Persekutuan Terkecil, Faktorial, Faktor Prima, Dll.

### Faktorial
Dalam matematika, faktorial bilangan bulat positif n, dilambangkan dengan n!, adalah hasil kali semua bilangan bulat positif kurang dari atau sama dengan n: 5! = 5 x 4 x 3 x 2 x 1 = 120

### Bilangan prima
Versi yang paling naif adalah menguji menurut definisi, yaitu menguji apakah N habis dibagi dengan pembagi € [2..N-1]. Ini berfungsi, tetapi berjalan di O(N). Peningkatan besar pertama adalah menguji apakah N habis dibagi oleh pembagi € [2..sqrt(N)]

### Faktor Persekutuan Terbesar (FPB)
Pembagi persekutuan terbesar dari bilangan bulat a dan b, dilambangkan fpb(a, b), adalah bilangan bulat terbesar yang membagi a dan b. 
Misalnya, fpb(30, 12)= 6. Algoritma Euclid menyediakan cara yang efisien untuk menghitung nilai fpb(a, b). Algoritma didasarkan pada rumus.

### Kelipatan Persekutuan Terkecil (KPK)
Kelipatan persekutuan terkecil, dilambangkan dengan Icm(a,b), yang merupakan bilangan bulat terkecil yang habis dibagi oleh a dan b


## Searching
Pencarian adalah proses menemukan posisi nilai tertentu dalam daftar nilai


## Sorting
Sorting adalah proses menyusun data dalam urutan tertentu. Biasanya, kami mengurutkan berdasarkan nilai elemen. Kita dapat mengurutkan angka, kata, pasangan, dll. Misalnya, kita dapat mengurutkan siswa berdasarkan tinggi badannya, dan kita dapat mengurutkan kota menurut abjad atau menurut jumlah warganya. Urutan yang paling sering digunakan adalah urutan angka dan urutan abjad.

### Beberapa contoh algoritma sorting :
#### Selection sort - O(n^2)
Idenya: Temukan elemen minimal dan tukar dengan elemen pertama dari array. Selanjutnya, urutkan sisa array, tanpa elemen pertama, dengan cara yang sama. 
Perhatikan bahwa setelah k iterasi (pengulangan semua yang ada di dalam loop) ke elemen pertama akan diurutkan dalam urutan yang benar (jenis properti ini disebut loop invarian)

#### Counting sort - O(n + k)
Idenya: Pertama, hitung elemen dalam larik penghitung. Selanjutnya, lakukan iterasi melalui array penghitung dalam urutan yang meningkat.
Perhatikan bahwa kita harus mengetahui rentang nilai yang diurutkan. Jika semua elemen berada di himpunan {0,1,..., k), maka larik yang digunakan untuk menghitung harus berukuran k + 1. Batasan di sini mungkin memori yang tersedia.

#### Merge sort - O (log n)
Idenya: Bagilah array yang tidak disortir menjadi dua bagian, urutkan setiap bagian secara terpisah dan kemudian gabungkan saja. Setelah dibelah, masing- masing bagian dibelah dua lagi. Kami mengulangi algoritme ini hingga kami mendapatkan elemen individual, yang diurutkan berdasarkan definisi. Penggabungan dua larik terurut yang terdiri dari k elemen memerlukan waktu O(k); cukup berulang kali memilih yang lebih rendah dari elemen pertama dari dua bagian yang digabungkan. tingkat berturut- turut Panjang array dibelah dua pada setiap iterasi. Dengan cara ini, kita mendapatkan potongan 1, 2, 4, 8, .... Untuk setiap level, penggabungan dari semua pasangan irisan yang berurutan membutuhkan O(n) waktu. Jumlah level adalah O(log n), sehingga total kompleksitas waktu adalah O(n log n).
