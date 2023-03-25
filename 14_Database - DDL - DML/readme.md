# (14) Database - DDL - DML

Database adalah sekumpulan data yang terorganisir. Database memiliki beberapa relationship, diantaranya :
- One To One : satu item dalam tabel diasosiasikan dengan satu item dalam tabel lain.
- One To Many : satu item dalam tabel yang berhubungan dengan banyak item dalam tabel lain.
- Many to Many : satu atau lebih item dalam satu tabel dapat memiliki hubungan dengan satu atau lebih item dalam tabel lain.

## RDBMS
Relational Database Management System (RDBMS) adalah kumpulan program dan kemampuan yang memungkinkan tim IT dan lainnya untuk membuat, memperbarui, mengelola, dan berinteraksi dengan basis data relasional. Contoh software yang menggunakan Relational Database Model sebagai dasarnya adalah MySQL.

## SQL
Structured Query Language (SQL) adalah bahasa standar untuk menyimpan, memanipulasi, dan mengambil data dalam basis data. SQL memiliki beberapa jenis perintah, diantaranya :
1. Data Definition Languge (DDL) : terdiri dari perintah SQL yang dapat digunakan untuk mendefinisikan skema database. Daftar perintah DDL:
    - CREATE: Perintah ini digunakan untuk membuat database atau objeknya (seperti tabel, indeks, fungsi, tampilan, prosedur penyimpanan, dan pemicu).
    - DROP: Perintah ini digunakan untuk menghapus objek dari database.
    - ALTER: Ini digunakan untuk mengubah struktur database.
    - TRUNCATE: Ini digunakan untuk menghapus semua catatan dari tabel, termasuk semua ruang yang dialokasikan untuk catatan dihapus.
    - RENAME: Ini digunakan untuk mengganti nama objek yang ada di database.

2. Data Manipulation Language (DML) : Perintah SQL yang berhubungan dengan manipulasi data yang ada dalam database. Daftar perintah DML:
    - INSERT : Digunakan untuk memasukkan data ke dalam tabel.
    - UPDATE: Digunakan untuk memperbarui data yang ada dalam tabel.
    - DELETE : Digunakan untuk menghapus record dari tabel database.

3. Data Control Language (DCL) : Perintah SQL yang berhubungan dengan hak, izin, dan kontrol lain dari sistem basis data. Daftar perintah DCL:
    - GRANT: Perintah ini memberi pengguna hak akses ke database.
    - REVOKE: Perintah ini menarik hak akses pengguna yang diberikan dengan menggunakan perintah GRANT.