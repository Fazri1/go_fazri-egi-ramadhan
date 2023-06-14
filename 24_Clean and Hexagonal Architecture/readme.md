# (24) Clean and Hexagonal Architecture

Arsitektur yang baik adalah pemisahan perhatian menggunakan layer untuk membangun aplikasi yang modular, dapat diskalakan, dapat dipelihara, dan dapat diuji.

Mengerjakan proyek pemrograman mirip dengan merencanakan dan membangun distrik perumahan. Jika Anda tahu bahwa distrik tersebut akan berkembang dalam waktu dekat, Anda perlu menyediakan ruang untuk perbaikan di masa mendatang. Tanpa itu, Anda akan terpaksa menghancurkan bangunan dan jalan untuk memberi ruang bagi bangunan baru

Setiap tim memiliki strukturnya sendiri
- Bagaimana mempertahankan ini?
- Masalah mobilitas → sulit untuk mentransfer pengetahuan pada domain bisnis dan ditambah dengan struktur kode yang unik, kami memiliki waktu dan otak ekstra untuk mempelajarinya dari atas ke bawah.
- Masalah lain → ex: sulit untuk mengimplementasikan unittest terutama tes yang memiliki koneksi dengan database dan kemudian kami hampir tidak memilih tes integrasi.

### History
Selama beberapa tahun terakhir kami telah melihat berbagai ide mengenai arsitektur sistem.
- Hexagonal Architecture
- Onion Architecture
- Screaming Architecture
- DCI from Agile Development
- BCE from Object Oriented Software
- Clean Architecture

Meskipun semua arsitektur ini agak berbeda dalam detailnya, mereka sangat mirip. Semuanya memiliki tujuan yang sama, yaitu pemisahan kepentingan. Mereka semua mencapai pemisahan ini dengan membagi perangkat lunak menjadi beberapa layer. Masing-masing memiliki setidaknya satu layer untuk aturan bisnis, dan satu lagi untuk antarmuka.

Adapun kendala sebelum merancang Clean Architecture adalah :
- Independen Kerangka
Ini memungkinkan Anda untuk menggunakan kerangka kerja seperti alat, daripada memiliki kendala terbatas.
- Dapat diuji
Aturan bisnis dapat diuji tanpa hal lain.
- Independen UI
UI dapat berubah dengan mudah, tanpa mengubah sistem lainnya.
- Independen dari Database
Aturan bisnis Anda tidak terikat dengan database, Anda dapat menukar database dengan mudah.
- Independen dari eksternal apapun
Faktanya aturan bisnis Anda sama sekali tidak tahu apa-apa tentang dunia luar

Keuntungan
- struktur standar, sehingga mudah menemukan jalan Anda dalam proyek,
- pengembangan lebih cepat dalam jangka panjang,
- ketergantungan mocking menjadi sepele dalam pengujian unit,
- peralihan yang mudah dari prototipe ke solusi yang tepat (mis., mengubah penyimpanan dalam memori ke database SQL).

CA Layer
Kami akan mendefinisikan arsitektur 3 layer klasik (kami dapat memiliki lebih banyak layer).
- (opsional) layer entitas - objek bisnis karena mencerminkan konsep yang dikelola aplikasi Anda
- Kasus Penggunaan - layer domain - berisi logika bisnis
- Pengontrol - layer presentasi - menampilkan data ke layar dan menangani interaksi pengguna
- Driver - layer data - Mengelola data aplikasi mis. mengambil data dari jaringan, mengelola cache data

## DDD (Domain Driven Design)
“Saya tidak tahu tentang topiknya dan saya ingin tahu agar dapat melakukan pekerjaan saya dengan baik”

Dua Pandangan
- Software Engineer
- Pakar dalam sesuatu yang ingin kami selesaikan (pemilik produk, manajer proyek, klien)
Mereka berbicara bahasa yang berbeda, bagaimana kita menggabungkan kedua perspektif?

Ubiquitous Languange
Singkat cerita setelah bertanya dengan ahli, eksplorasi, menemukan domain pekerjaan, modeling, dan brainstorming akhirnya kami memiliki kesamaan dalam perspektif.

You need to go slow, to go fast
Pastikan Anda memecahkan masalah yang valid dengan cara yang optimal. Setelah itu terapkan solusi dengan cara yang akan dipahami bisnis Anda tanpa terjemahan tambahan dari bahasa teknis yang diperlukan.

Clean Architecture adalah arsitektur perangkat lunak
Domain Driven Design adalah teknik desain perangkat lunak