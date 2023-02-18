// Prioritas 1-1
// menghitung luas trapesium

package main

import "fmt"

func main() {
	var sisiAtas, sisiBawah, tinggi float32
	
	fmt.Print("Sisi atas: ")
	fmt.Scan(&sisiAtas)
	fmt.Print("Sisi bawah: ")
	fmt.Scan(&sisiBawah)
	fmt.Print("Tinggi: ")
	fmt.Scan(&tinggi)

	luas := 0.5 * (sisiAtas + sisiBawah) * tinggi
	fmt.Println("Luasnya adalah", luas, "cm")
}
