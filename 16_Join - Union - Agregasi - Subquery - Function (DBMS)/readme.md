# (16) Join - Union - Agregasi - Subquery - Function (DBMS)

## JOIN
Sebuah klausa untuk mengkombinasikan record dari dua atau lebih tabel

JOIN STANDAR SQL ANSI
1. INNER JOIN
Inner join akan mengembalikan baris-baris dari dua tabel atau lebih yang memenuhi syarat.

2. LEFT JOIN
Left join akan mengembalikan seluruh baris dari tabel disebelah kiri yang dikenai kondisi ON dan hanya baris dari tabel disebelah kanan yang memenuhi kondisi join.

3. RIGHT JOIN
Right join akan mengembalikan semua baris dari tabel sebelah kanan yang dikenai kondisi ON dengan data dari tabel sebelah kiri yang memenuhi kondisi join. Teknik ini merupakan kebalikan dari left join.

## UNION
Operator UNION digunakan untuk menggabungkan result-set dari dua atau lebih pernyataan SELECT. Ada hal yang perlu diperhatikan dari union adalah jumlah field yang dikeluarkan/dipanggil harus sama.

## AGGREGATE
Fungsi agregasi adalah fungsi di mana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal.

FUNGSI AGREGASI SQL
- MIN
fungsi di mana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal.

- MAX
Digunakan untuk mendapatkan nilai maximum atau nilai terbesar dari sebuah data record di tabel.

- SUM
Digunakan untuk mendapatkan jumlah total nilai dari sebuah data atau record di tabel.

- AVG
Digunakan untuk mencari nilai rata-rata (average) dari sebuah data atau record di tabel.

- COUNT
Digunakan untuk mencari jumlah dari sebuah data atau record di tabel.

- HAVING
Digunakan untuk menyeleksi data berdasarkan kriteria tertentu, dimana kriteria berupa fungsi aggregat.

## SUBQUERY
Subquery atau Inner query atau Nested query adalah query di dalam query SQL lain. Sebuah subquery digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil. Subquery dapat digunakan dengan SELECT, INSERT, UPDATE, dan DELETE statements bersama dengan operator seperti

Peraturan yang harus ditaati
- Harus tertutup dalam tanda kurung.
- Sebuah subquery hanya dapat memiliki satu kolom pada klausa SELECT, kecuali beberapa kolom yang di query utama untuk subquery untuk membandingkan kolom yang dipilih.
- Subqueries yang kembali lebih dari satu baris hanya dapat digunakan dengan beberapa value operator, seperti operator IN.
- Daftar SELECT tidak bisa menyertakan referensi ke nilai-nilai yang mengevaluasi ke BLOB, ARRAY, CLOB, atau NCLOB.
- Sebuah subquery tidak dapat segera tertutup dalam fungsi set.

## FUNCTION
sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya

CONTOH
Function untuk mengembalikan jumlah data dari tweets per user
``` sql
DELIMITER $$
CREATE FUNCTION sf_count_tweet_peruser
(user_id_p int) RETURNS INT DETERMINISTIC
BEGIN
DECLARE total INT;
SELECT COUNT(*) INTO total FROM tweets
WHERE user_id = user_id_p AND
type='tweets';
RETURN total;
END$$
DELIMITER ;
```

PENJELASAN
DELIMITER,
memberi tahu kepada mysql soal delimiter yang digunakan, secara default menggunakan : jadi bila ada tanda; mysql akan mengartikan akhir dari statement, pada contoh di atas delimeter yang digunakan $$ jadi akhir statementnya adalah $$

CREATE FUNCTION, 
adalah header untuk membuat function

RETURN,
adalah untuk menentukan tipe data yang di return-kan oleh function

DETERMINISTIC/NOT DETERMINISTIC,
adalah untuk menentukan yang bisa menggunakan function ini adalah user pembuatnya saja (determinisric) atau user siapa saja (not determinisric)

BEGIN END,
adalah body dari function jadi semua SQL nya di tulis disini

CONTOH
Buat Trigger Function untuk delete data yang berhubungan dengan table users yang sedang di delete datanya
``` sql
DELIMITER $$
CREATE TRIGGER delete_all_data_user
BEFORE DELETE ON users FOR EACH ROW
BEGIN
-- declare variables
DECLARE v_user_id INT;
SET v_user_id=OLD.id;

-- trigger code
DELETE FROM tweets WHERE user_id=v_user_id;
DELETE FROM user _ followers WHERE user_id=v_user_id;
END$$
DELIMITER ;
```

PENJELASAN
DELIMITER,
adalah untuk memberi tahu kepada myql soal delimiter yang digunakan, secara default menggunakan jadi bila ada tanda; mysql akan mengartikan akhir dari statement, pada contoh di atas delimeter yang digunakan $$ jadi akhir statementnya adalah $$

CREATE TRIGGER,
adalah header untuk membuat trigger function

DECLARE,
adalah syntax untuk mendeclare kan sesuatu hal/variabel

OLD,
adalah variabel yang menyimpan nilai dari dalam data yang sedang berinteraksi/dipanggil

NEW, 
adalah variabel yang menyimpan nilai dari data yang baru masuk/sedang di input

BEGIN END,
adalah body dari function jadi semua SQL nya di tulis disini