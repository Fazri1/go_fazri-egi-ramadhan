# (18) System Design

## Diagram
Diagram adalah representasi simbolis dari informasi menggunakan teknik visualisasi.

### Common used software design
- Flowchart
Teori - mewakili 1 proses. Proses, Keputusan, Terminator.

- Use Case Diagram
Diagram use case meringkas detail pengguna sistem Anda (juga dikenal sebagai aktor) dan interaksi mereka dengan sistem.

- Entity Relationship Diagram
Diagram Hubungan Entitas (ER) adalah jenis diagram alur yang menggambarkan bagaimana "entitas" seperti orang, objek, atau konsep berhubungan satu sama lain dalam suatu sistem.

- High Level Architecture
Desain sistem secara keseluruhan. Periksa audiens target Anda.

## Horizontal scaling and vertical scaling
Dasar-dasar Desain Sistem
Setiap kali kami merancang sistem yang besar, kami perlu mempertimbangkan beberapa hal:
1. Apa saja potongan arsitektur berbeda yang dapat digunakan?
2. Bagaimana potongan-potongan ini bekerja satu sama lain?
3. Bagaimana kita dapat memanfaatkan bagian-bagian ini dengan sebaik-baiknya: apa pengorbanan yang tepat?
Mengakrabkan konsep-konsep ini akan sangat bermanfaat dalam memahami konsep sistem terdistribusi.

## Key Characteristics of Distributed Systems
#### Scalability
Skalabilitas adalah kemampuan sistem, proses, atau jaringan untuk tumbuh dan mengelola peningkatan permintaan. Sistem terdistribusi apa pun yang dapat terus berkembang untuk mendukung jumlah pekerjaan yang terus bertambah dianggap dapat diskalakan.
Suatu sistem mungkin harus diskalakan karena banyak alasan seperti peningkatan volume data atau peningkatan jumlah pekerjaan, misalnya jumlah transaksi. Sistem yang dapat diskalakan ingin mencapai penskalaan ini tanpa kehilangan kinerja.

Vertical Scale example : system upgrade
Horizontal Scale : system duplication

Yang mana yang Anda sukai? Mengapa?
Jika kita masih memiliki kemungkinan untuk meningkatkan kemampuan dan dapat menjamin itu akan bertahan untuk jangka panjang maka skala vertikal adalah kandidat yang baik
→ tidak ada perubahan pada kode.

Jika kita tidak dapat memprediksi dan ada kemungkinan kenaikan kita akan terjadi dengan cepat, maka aman untuk mengikuti pendekatan skala horizontal.
→ ada kemungkinan ada perubahan pada kode untuk diadaptasi.

#### Reliability
Menurut definisi, keandalan adalah probabilitas sistem akan gagal dalam periode tertentu. Secara sederhana, sistem terdistribusi dianggap andal jika tetap memberikan layanannya bahkan ketika satu atau beberapa komponen perangkat lunak atau perangkat kerasnya gagal. Keandalan merupakan salah satu karakteristik utama dari setiap sistem terdistribusi, karena dalam sistem seperti itu setiap mesin yang gagal selalu dapat diganti dengan mesin lain yang sehat, memastikan penyelesaian tugas yang diminta.

#### Availability
Menurut definisi, ketersediaan adalah waktu sistem tetap beroperasi untuk melakukan fungsi yang diperlukan dalam periode tertentu. Ini adalah ukuran sederhana persentase waktu sistem, layanan, atau mesin tetap beroperasi dalam kondisi normal.
Contoh: pesawat yang dapat diterbangkan berjam-jam dalam sebulan tanpa banyak downtime dapat dikatakan memiliki ketersediaan yang tinggi.

#### Efficiency
Untuk memahami cara mengukur efisiensi sistem terdistribusi, anggaplah kita memiliki operasi yang berjalan secara terdistribusi dan mengirimkan sekumpulan item sebagai hasilnya. Dua ukuran standar efisiensinya adalah waktu respons (atau latensi) yang menunjukkan penundaan untuk mendapatkan item pertama dan throughput (atau bandwidth) yang menunjukkan jumlah item yang dikirimkan dalam satuan waktu tertentu (misalnya, satu detik).

#### Serviceability or Manageability
Serviceability atau pengelolaan adalah kesederhanaan dan kecepatan dimana sistem dapat diperbaiki atau dipelihara; jika waktu untuk memperbaiki sistem yang gagal bertambah, maka ketersediaannya akan berkurang. Hal-hal yang perlu dipertimbangkan untuk keterkelolaan adalah kemudahan mendiagnosis dan memahami masalah saat terjadi, kemudahan melakukan pembaruan atau modifikasi, dan seberapa sederhana pengoperasian sistem (yaitu, apakah sistem beroperasi secara rutin tanpa kegagalan atau pengecualian?).


## Job/Work Queue
Dalam perangkat lunak sistem, antrian pekerjaan (terkadang antrian batch), adalah struktur data yang dikelola oleh perangkat lunak penjadwal pekerjaan yang berisi pekerjaan untuk dijalankan.

Work Queue adalah framework untuk membangun aplikasi master-worker besar yang menjangkau ribuan mesin yang diambil dari cluster, cloud, dan grid.

## Load Balancing
Load Balancer (LB) adalah komponen penting lainnya dari sistem terdistribusi apa pun. Ini membantu untuk menyebarkan lalu lintas di sekelompok server untuk meningkatkan daya tanggap dan ketersediaan aplikasi, situs web, atau basis data. LB juga melacak status semua sumber daya sambil mendistribusikan permintaan. Jika server tidak tersedia untuk menerima permintaan baru atau tidak merespons atau memiliki tingkat kesalahan yang tinggi, LB akan berhenti mengirimkan lalu lintas ke server tersebut.

Untuk memanfaatkan skalabilitas dan redundansi penuh, kami dapat mencoba menyeimbangkan beban di setiap lapisan sistem. Kita dapat menambahkan LB di tiga tempat:
1. Antara pengguna dan server web
2. Antara server web dan lapisan platform internal, seperti server aplikasi atau server cache
3. Antara lapisan platform internal dan basis data.

## Monolithic and Microservices
1. Monolitik
Aplikasi monolitik memiliki basis kode tunggal dengan banyak modul. Modul dibagi menjadi fitur bisnis atau fitur teknis. Ini memiliki sistem build tunggal yang membangun seluruh aplikasi dan/atau ketergantungan. Ini juga memiliki biner tunggal yang dapat dieksekusi atau digunakan.

2. Layanan mikro
Layanan mikro adalah layanan yang dapat diterapkan secara independen yang dimodelkan di sekitar domain bisnis. Mereka berkomunikasi satu sama lain melalui jaringan, dan sebagai pilihan arsitektur menawarkan banyak pilihan untuk memecahkan masalah yang mungkin Anda hadapi. Oleh karena itu, arsitektur layanan mikro didasarkan pada beberapa layanan mikro yang berkolaborasi.

## SQL and NoSQL
SQL dan NoSQL (atau database relasional dan database non-relasional).
- Database relasional terstruktur dan memiliki skema standar seperti buku telepon yang menyimpan nomor telepon dan alamat.
- Basis data non-relasional tidak terstruktur, dan memiliki skema dinamis seperti folder file yang menyimpan segala sesuatu mulai dari alamat dan nomor telepon seseorang hingga 'suka' Facebook dan preferensi belanja online mereka.

### SQL
BENEFIT RELATIONAL DB
- Dirancang untuk segala keperluan
- SQL - Memiliki standar yang jelas
- Memiliki banyak tool (administrasi, reporting, framework)

ACID
- Atomicity; transaksi terjadi semua atau tidak sama sekali
- Consistency; data tertulis merupakan data valid yang ditentukan berdasarkan aturan tertentu
- Isolation; pada saat terjadi request yang bersamaan (concurrent), memastikan bahwa transaksi dieksekusi seperti dijalankan secara sekuensial
- Durability; jaminan bahwa transaksi yang telah tersimpan, tetap tersimpan.

### Not only SQL
adalah cara berpikir yang benar-benar baru tentang database. NoSQL bukan database relasional.

DBMS yang menyediakan mekanisme yang lebih fleksibel dibandingkan dengan model RDBMS (Sifat ACID).
Menghindari:
- Effort pada sifat transaksi ACID
- Kompleksitas SQL
- Design schema di depan
- Transactions (ditangani oleh aplikasi)

Kelebihan: schema less, fast development, etc.
Kapan digunakan: membutuhkan skema fleksibel, ACID tidak diperlukan, data logging (json), data sementara (chace), etc
Kapan tidak direkomendasikan : data finansial, data transaksi, business critical, ACID - compliant.

#### Schema-less
1. Traditional RDBMS
- Tidak bisa menambah data yang tidak sesuai skema
- Perlu menambahkan data NULL pada item yang tidak memiliki data
- Memiliki tipe data yang strict
- Tidak dapat menambah beberapa item data pada sebuah field

2. NoSQL DBMS
- Tidak memiliki skema ketika menambahkan data
- Aplikasi menangani proses validasi tipe data
- Mendukung proses aggregasi dokumen pada item.


Dalam hal teknologi basis data, tidak ada solusi yang cocok untuk semua. Itu sebabnya banyak bisnis mengandalkan database relasional dan non-relasional untuk kebutuhan yang berbeda. Meskipun database NoSQL semakin populer karena kecepatan dan skalabilitasnya, masih ada situasi di mana database SQL yang sangat terstruktur dapat bekerja lebih baik; memilih teknologi yang tepat bergantung pada use case.

## Caching
Cache yang digunakan dalam data yang baru-baru ini diminta kemungkinan besar akan diminta lagi. Mereka digunakan di hampir setiap lapisan komputasi: perangkat keras, sistem operasi, browser web, aplikasi web, dan banyak lagi.
Cache seperti memori jangka pendek: memiliki jumlah ruang terbatas, tetapi biasanya lebih cepat daripada sumber data asli dan berisi item yang paling baru diakses. Cache dapat ada di semua level dalam arsitektur, tetapi sering ditemukan di level yang paling dekat dengan ujung depan tempat cache diimplementasikan untuk mengembalikan data dengan cepat tanpa membebani level downstream.

## Database Replication
Redundancy and Replication
Redundansi adalah duplikasi komponen atau fungsi penting dari suatu sistem dengan maksud untuk meningkatkan keandalan sistem, biasanya dalam bentuk cadangan atau fail-safe, atau untuk meningkatkan kinerja sistem yang sebenarnya. Misalnya, jika hanya ada satu salinan file yang disimpan di satu server, maka kehilangan server itu berarti kehilangan file tersebut. Karena kehilangan data jarang merupakan hal yang baik, kami dapat membuat salinan file duplikat atau redundan untuk mengatasi masalah ini.

## Database Indexing
Pengindeksan adalah cara untuk mengoptimalkan kinerja database dengan meminimalkan jumlah akses disk yang diperlukan saat kueri diproses. Ini adalah teknik struktur data yang digunakan untuk menemukan dan mengakses data dalam database dengan cepat.

Sebagian besar database menggunakan B-Tree sebagai struktur data mereka untuk pengindeksan. Dan B-Tree memiliki kompleksitas O(log n) untuk operasi pencarian, penghapusan dan penyisipan.

Indeks adalah bagian terpisah yang menyimpan nilai kolom + penunjuk yang diindeks ke baris yang diindeks dalam urutan yang diurutkan untuk pencarian yang efisien.