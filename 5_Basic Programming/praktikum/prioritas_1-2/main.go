// Prioritas 1-2
// menentukan apakah sebuah bilang adalah bilang ganjil atau genap

package main

import "fmt"

func main() {
	var angka int
	
	fmt.Print("Angka: ")
	fmt.Scan(&angka)

	if angka % 2 == 0 {
		fmt.Println(angka, "adalah bilangan genap")
	} else {
		fmt.Println(angka, "adalah bilangan ganjil")
	}
}