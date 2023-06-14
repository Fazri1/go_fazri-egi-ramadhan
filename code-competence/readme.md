# Code Competence
## Scenario
Sebuah toko barang elektronik ingin merancang sebuah aplikasi yang dapat digunakan untuk mengelola data barang elektronik. Data barang terdiri dari ID, nama, deskripsi, jumlah Stok, harga dan kategori barang.

### Endpoint
| URI | Method | Description |
| --- | :---: | --- |
| /items | GET | Mendapatkan semua data barang |
| /items/:id | GET | Mendapatkan data barang berdasarkan ID |
| /items | POST | Menambahkan data barang |
| /items/:id | PUT | Mengubah data barang |
| /items/:id | DELETE | Menghapus data barang |
| /items/category/:category_id | GET | Mendapatkan semua data barang berdasarkan ID kategori |
| /items?keyword=item_name | GET | Mendapatkan data barang berdasarkan nama barang |
| /register | POST | Membuat akun atau registrasi |
| /login | POST | Masuk dengan akun yang sudah dibuat |

Aplikasi sudah menerapkan :
- Menerapkan logging middleware.
- Menerapkan authentication middleware dengan menggunakan token JWT. Middleware ini digunakan untuk semua endpoint diatas, kecuali register dan login.
- Aplikasi menggunakan penyimpanan persisten dengan menggunakan MySQL dengan menggunakan library GORM.
