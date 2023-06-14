# (19) Intro RESTful API

## API
Application Programming Interface (API) adalah sekumpulan fungsi dan prosedur yang memungkinkan pembuatan aplikasi yang mengakses fitur atau data dari sistem operasi, aplikasi, atau layanan lainnya.

## REST API
Representational State Transfer API adalah API yang sesuai dengan prinsip desain REST, atau gaya arsitektur transfer status representasional.

Best Practice dalam mendesain sebuah API :
1. Gunakan kata "benda" dibandingkan dengan "kata kerja" karena method sudah mewakili kata kerja. Contoh `/books/123`, bukan `/addBook123`.
2. Penggunaan kata "Jamak" atau "tuggal". Contoh `/cars/123`, bukan `/car/123`.
3. Gunakan Resource Nesting untuk menampilkan relasi atau hierarki. Contoh :
```
/users <- user's list
/users/123 <- specific user
/users/123/orders <- orders list that belongs to a specific user
/users/123/orders/0001 <- specific order of a specific user
```
4. Usahakan untuk bisa menghandle Trailing Slashes. Contoh `POST: /cars`, bukan `POST: /cars/`.

## Package Go for Build API
1. net/http : Menyediakan implementasi klien dan server http.
2. encoding/json : BERISI BANYAK FUNGSI UNTUK OPERASI JSON:
- decode JSON ke struct object
- decode JSON ke map[string]interface{} & interface{}
- decode array JSON ke array object
- encode object ke JSON string