// SOAL
// 1. Berapa banyak kekurangan dalam penulisan kode tersebut?
// 2. Bagian mana saja terjadi kekurangan tersebut?
// 3. Tuliskan alasan dari setiap kekurangan tersebut.

// JAWAB
// 1. terdapat 7 kekurangan

/* 2. 
1 penamaan struct
2 struct userservice
3 variabel t
4 variabel u
5 method getallusers
6 method getuserbyid
7 variabel r
*/

// 3.Alasan diberikan dalam bentuk komentar pada kode yang disediakan berikut.
package main

// kekurangan dalam penamaan struct (lebih baik diawali dengan huruf kapital)
type user struct {
	id       int
	username int
	password int
}

// - kekurangan dalam penamaan struct (lebih baik diawali dengan huruf kapital)
// - kekurangan dalam penamaan dua kata
type userservice struct {
	// kekurangan dalam penamaan variabel t
	t []user
}

// - kekurangan dalam penamaan alias untuk struct userservice
// - kekurangan dalam penamaan dua kata
func (u userservice) getallusers() []user {
	return u.t
}

// kekurangan dalam penamaan alias untuk struct userservice
// kekurangan dalam penamaan dua kata
func (u userservice) getuserbyid(id int) user {
	// kekurangan dalam penamaan variabel r
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}