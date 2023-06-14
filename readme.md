# (27) Compute Service

System & Software Deployment
Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen. Penyebarannya dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web & api ke server sedangkan aplikasi mobile ke Playstore/Appstore.

Strategi Deployment
- Big-Bang Deployment Strategy atau sering disebut Replace Deployment Strategy
- Rollout Deployment Strategy 
- Blue/Green Deployment Strategy
- Canary Deployment Strategy

Big-Bang/Replace Deployment Strategy
Kelebihan:
- Mudah di-implementasikan. Cara klasik, tinggal replace.
- Perubahan kepada sistem langsung 100% secara instan.

Kekurangan:
- Terlalu beresiko, rata-rata downtime cukup lama.

Rollout Deployment Strategy
Kelebihan:
- Lebih aman dan less downtime dari versi sebelumnya

Kekurangan:
- Akan ada 2 versi aplikasi berjalan secara barengan sampai semua server terdeploy, dan bisa membuat bingung. 
- Karena sifatnya perlahan satu persatu, untuk deployment dan rollback lebih lama dari yang Bigbang, karena prosesnya perlahan-lahan sampai semua server terkena efeknya.
- Tidak ada kontrol request. Server yang baru ter-deploy dengan aplikasi versi baru, langsung mendapat request yang sama banyaknya dengan server yang lain.

Blue/Green Deployment Strategy
Kelebihan:
- Perubahan sangat cepat, sekali switch service langsung berubah 100%.
- Tidak ada issue beda versi pada service seperti yang terjadi pada Rollout Deployment.

Kekurangan:
- Resource yang dibutuhkan lebih banyak. Karena untuk setiap deployment kita harus menyediakan service yang serupa environmentnya dengan yang sedang berjalan di production. 
- Testing harus benar-benar sangat diprioritaskan sebelum di switch, aplikasi harus kita pastikan aman dari request yang tiba-tiba banyak.

Canary Deployment Strategy
Kelebihan:
- Cukup aman
- Mudah untuk rollback jika terjadi error/bug, tanpa berimbas kesemua user

Kekurangan:
- Untuk mencapai 100% cukup lama dibanding dengan Blue/ Green deployment. Dengan Blue/Green deployment, aplikasi langsung 100% terdeploy keseluruh user.
