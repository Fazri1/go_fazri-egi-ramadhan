# (12) Concurrency in Golang

## Sequential VS Parallel VS Concurrent
### Sequential Program
Dalam program sequential, sebelum tugas baru dimulai, tugas sebelumnya harus selesai.

### Parallel Program
Dalam program paralel, banyak tugas dapat dijalankan secara bersamaan. (membutuhkan mesin multi-core)

### Concurrent Program
Dalam program concurrent, banyak tugas dapat dijalankan secara independen dan mungkin muncul secara bersamaan. Paralel artinya serentak, sedangkan concurrent artinya mandiri.

Concurrency mengizinkan parallelisme 
Konkurensi Go (goroutine) memudahkan pembuatan program paralel yang memanfaatkan mesin dengan banyak prosesor (kebanyakan mesin saat ini)

## Goroutine
Beberapa fitur Go :
- Eksekusi serentak (Goroutine)
- Sinkronisasi dan pesan (Channel)
- Kontrol serentak multi- arah (Select)

Goroutine adalah fungsi atau metode yang dijalankan secara bersamaan (independen) dengan fungsi atau metode lain. 
Biaya pembuatan Goroutine kecil jika dibandingkan dengan thread. thread adalah proses yang ringan, atau dengan kata lain thread adalah unit yang mengeksekusi kode di bawah program.

### Gomaxprocs
Gomaxprocs digunakan untuk mengontrol jumlah thread sistem operasi yang tersedia untuk eksekusi Goroutine ke program go tertentu.

## Channel & Select
### Channel
Channel adalah objek komunikasi yang digunakan goroutine untuk berkomunikasi satu sama lain. Channel terbagi menjadi 2, yaitu :
- Unbuffered Channel
`c := make(chan string)`

- Buffered Channel
`c := make(chan string, 42)`

Dalam menggunakan unbuffered channel diperlukan Goroutine, tapi buffered channel tidak perlu

### Select
Select membuatnya lebih mudah untuk mengontrol komunikasi data melalui satu atau banyak channel.


## Race Condition
Race condition adalah dimana 2 thread mengakses memori secara bersamaan, salah satunya adalah menulis. Race condition terjadi karena akses yang tidak sinkron ke memori bersama.

Terdapat beberapa cara untuk memperbaiki data race, diantaranya :
- WaitGroups
- Channels Blocking
- Mutex

### Blocking with WaitGroups
Cara paling mudah untuk menyelesaikan data race, adalah memblokir akses baca hingga operasi tulis selesai.

### Channel Blocking
Ini memungkinkan goroutine kita untuk melakukan sinkronisasi satu sama lain untuk sesaat.

### Mutex
Jika kami tidak peduli dengan urutan baca dan tulis, kami hanya meminta agar urutan tersebut tidak terjadi secara bersamaan. 
Jika ini terdengar seperti kasus penggunaan Anda, maka kami harus mempertimbangkan untuk menggunakan mutex