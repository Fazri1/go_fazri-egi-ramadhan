# (9) String, Advance Function, Pointer, Struct, Method,  Interface, Package & Error Handling

## String
Beberapa built-in function untuk manipulasi string pada golang, yaitu :
- Len -> Menghitung panjang string.
- Compare -> membandingkan dua string.
- Contains -> Untuk cek substr ada di dalam s.
- Replace -> Mengganti isi dari string.

## Function
### Variadic Function
Variadic function digunakan :
- Untuk melewatkan pembuatan slice sementara hanya untuk meneruskan ke fungsi
- Bila jumlah parameter input tidak diketahui
- Untuk mengungkapkan niat Anda untuk meningkatkan keterbacaan

### Anonymous function == literal function
Fungsi anonim adalah fungsi yang tidak mengandung nama apa pun. Ini berguna saat Anda ingin membuat fungsi sebaris.

### Closure
Closure adalah tipe khusus dari fungsi anonim yang mereferensikan variabel yang dideklarasikan di luar fungsi itu sendiri. Dalam hal ini kita akan menggunakan variabel yang tidak diteruskan ke fungsi sebagai parameter, melainkan tersedia saat fungsi dideklarasikan.

### Defer Function
Function yang hanya dijalankan setelah fungsi induknya dikembalikan. Pengembalian berganda juga dapat digunakan, mereka dijalankan sebagai tumpukan, satu per satu.


## *Pointer
Pointer adalah variabel yang menyimpan alamat memori dari variabel lain. Pointer memiliki kekuatan untuk mengubah data yang mereka tunjuk. Perubahan variabel dengan referensi memori yang sama akan mempengaruhi satu sama lain.

### Operator di Pointer
Pointer memiliki 2 operator penting, yaitu :
- /* Operator Dereferencing : Mendeklarasikan variabel penunjuk, Mengakses nilai yang tersimpan di alamat.
- & Operator Referencing : Mengembalikan alamat variabel, Mengakses alamat variabel ke pointer.


## Struct (Object di Golang)
Sebuah struct adalah tipe yang ditentukan pengguna yang berisi kumpulan bidang/properti atau fungsi (metode) bernama.

## Method
Method adalah fungsi yang melekat pada suatu tipe (bisa berupa struct atau tipe data lainnya).

### Deklarasi method
Sama seperti fungsi, hanya deklarasi variabel objek yang perlu ditambahkan di antara kata kunci func dan nama fungsi.
``` go
func (receiver StructType) MethodName(parameterList) (returnTypes) {
    // block statement
}
```

### Alasan menggunakan method daripada function
- Membantu Anda menulis kode gaya berorientasi objek di Go.
- Metode membantu Anda menghindari konflik penamaan.
- Pemanggilan metode jauh lebih mudah dibaca dan dipahami daripada pemanggilan fungsi.

## Interface
Interface adalah kumpulan tanda tangan metode yang dapat diimplementasikan objek. Karenanya antarmuka mendefinisikan perilaku objek.

## Package
Paket adalah kumpulan fungsi dan data.

Cara membuat variabel dapat diakses ke package lain :
- `var ageOfUniverse int` : Paket lain tidak dapat melihat ini karena huruf pertama adalah huruf kecil
- `var AgeOfUniverse int` : Paket lain dapat melihat ini karena huruf pertamanya adalah UPPERCASE

## Error Handling
ERROR, PANIC & RECOVER
tidak ada try catch di Golang, biasanya menggunakan Panic & Recover untuk penanganan error.

### Error handling object
Jika Anda menulis metode sendiri yang mengharuskan untuk mengembalikan kesalahan jika terjadi kesalahan di antaranya, gunakan paket 'error' untuk tujuan tersebut. 

### Panic 
Saat runtime Go mendeteksi kesalahan ini, ia akan panik.

### Recover
Untuk menambahkan kemampuan pemulihan dari kesalahan panik, tambahkan fungsi anonim atau tentukan fungsi khusus dan panggil dengan kata kunci 'defer' dari dalam metode, di mana kepanikan mungkin terjadi dari panggilan internal lainnya.