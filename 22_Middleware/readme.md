# (22) Middleware

Middleware
Middleware adalah entitas yang menghubungkan ke pemrosesan permintaan/respons server. Middleware, izinkan kami untuk menerapkan berbagai fungsi sekitar permintaan http masuk ke server anda dan respon keluar.

Example Third Party Middleware
- Negroni
- Echo
- Interpose
- Alice

### SETUP MIDDLEWARE ECHO
Echo#Pre()
Dieksekusi sebelum router memproses permintaan
- HTTPSRedirect
- HTTPSWWWRedirect
- WWWRedirect 
- AddTrailingSlash
- Remove TrailingSlash
- MethodOverride
- Rewrite

Echo#Use()
Dieksekusi setelah router memproses permintaan dan memiliki akses penuh ke echo.Context API
- BodyLimit
- Logger
- Gzip
- Recover
- BasicAuth
- JWTAuth
- Secure
- CORS
- Static

### Log Middleware
Logger Middleware
- Mencatat permintaan HTTP informasi.
- Sebagai footprint.
- Sumber data untuk analitik.

### Auth Middleware
Alasan menggunakan autentikasi:
- Identifikasi Pengguna
- Mengamankan Data dan Informasi

Contoh Authentication Middleware:
- Basic Authentication
- JSON Web Token

1. Basic Authentication
Basic auth adalah permintaan teknik otentikasi http, metode ini membutuhkan informasi nama pengguna dan kata sandi untuk dimasukkan ke dalam header permintaan.

Information Encoded Format :
```
'Authorization: Basic +
base64 encode('username: password')
Header Generate:
Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=
```

2. JWT Middleware
JSON Web Token (JWT) adalah sarana aman URL yang ringkas untuk mewakili klaim yang akan ditransfer antara dua pihak.
Information Encoded Format :
```
'Authorization: Bearer + Token
Header Generate :
Authorization: Bearer
eyJhbGci0iJIUZI1NiIsInR5cCI6IkpXVCJ9 eyJhdXRob3JpemVkIjp0cnVILCJ1eHA10jE2MTQwNzU3NDMsInVzZXJJZCI6MX0.3RGYB17pwh7J86Cmhcp78AMi3LU-10_Wuu5Rt098vCk
```