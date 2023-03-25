// kode yang sudah diperbaiki
package main

type Kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (mobil *Mobil) berjalan() {
	mobil.tambahKecepatan(10)
}

func (mobil *Mobil) tambahKecepatan(kecepatanBaru int) {
	mobil.kecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := Mobil{}
	mobilLamban.berjalan()
}