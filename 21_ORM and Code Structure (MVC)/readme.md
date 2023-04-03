# (21) ORM dan Struktur Kode (MVC)

## ORM (Object Relational Mapping)
Pemetaan objek-relasional (alat pemetaan ORM, O/RM, dan O/R) dalam ilmu komputer adalah teknik pemrograman untuk mengubah data antara sistem tipe yang tidak kompatibel menggunakan bahasa pemrograman berorientasi objek.

#### Keuntungan ORM
- Kueri yang kurang berulang
- Secara otomatis mengambil data ke objek siap pakai
- Cara sederhana jika Anda ingin menyaring data sebelum menyimpannya di database
- Beberapa memiliki fitur kueri cache 

#### Kekurangan ORM
- Tambahkan lapisan dalam kode dan biaya proses overhead
- Memuat data hubungan yang tidak diperlukan
- Permintaan mentah yang kompleks bisa lama ditulis dengan ORM (> 10 tabel bergabung)
- Fungsi SQL tertentu yang terkait dengan satu vendor mungkin tidak didukung atau tidak ada fungsi khusus (MySQL : FORCE INDEX)

### Database Migration
Cara memperbarui versi basis data untuk menyesuaikan perubahan versi aplikasi. Perubahan dapat berupa upgrade ke versi terbaru atau rollback ke versi sebelumnya.

Alasan menggunakan DB Migration
- Kesederhanaan pembaruan basis data
- Kesederhanaan rollback basis data
- Lacak perubahan pada struktur database
- Sejarah struktur basis data ditulis pada sebuah kode
- Selalu kompatibel dengan perubahan versi aplikasi

### GORM
Library ORM yang fantastis untuk Golang.

## Code Structure
Susun proyek Anda menggunakan Model-View-Controller (MVC).

MVC
MVC adalah kependekan dari Model, View, dan Controller. MVC adalah cara populer untuk mengatur kode Anda. Ide besar di balik MVC adalah bahwa setiap bagian dari kode Anda memiliki tujuan, dan tujuan tersebut berbeda.

Alasan menggunakan struktur
- Untuk mencapai aplikasi modular.
- Terapkan separation of concerns.
- Lebih sedikit konflik pada pembuatan versi.
