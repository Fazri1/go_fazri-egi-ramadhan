# (26) Docker

### Container 
Container bukanlah mesin virtual.
Jawaban Singkat: Container adalah proses ... dengan isolasi sistem file.
Jawaban Panjang: Semua yang ada di Linux adalah file.
* /dev/sda = hard disk
* /dev/proc = processes
* /dev/usb
* /dev/cpu
* /dev/std(in/out)
* /bin/bash... just a binary file

### Container vs Virtual Machines
Container
- Abstraksi pada lapisan aplikasi
- Wadah mengambil lebih sedikit ruang daripada VM (gambar wadah biasanya berukuran puluhan MBS)
- Menangani lebih banyak aplikasi dan membutuhkan lebih sedikit VM dan sistem Operasi.

Virtual Machines
- Abstraksi perangkat keras fisik.
- Setiap VM menyertakan salinan lengkap dari sistem operasi
- Juga lambat untuk boot.

### Docker Volume
Docker Volume bisa dianggap sebagai storage atau tempat penyimpanan data di container, Tentunya saat kita membuat container kita tidak ingin ketika container kita mati data yang ada pada container ikut terhapus juga. Untuk itu kita dapat memanfaatkan Volume pada docker.
Sebelum kita membuat volume kita perlu tahu dimana container menyimpan data. Kita bisa mengetahuinya melalui docker hub, contoh mysql menyimpan data pada /my/own/datadir.

### Docker Network
Defaultnya container pada docker akan saling terisolasi satu sama lain. Kita tidak dapat melakukan request api (misal) dari container satu ke container Iain. Untuk itu kita harus membuat dan mendaftarkan container pada network yang sama.