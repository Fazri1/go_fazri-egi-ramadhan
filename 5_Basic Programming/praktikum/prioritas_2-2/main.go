// prioritas 2-2
// mencetak faktor bilangan dari sebuah angka
package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Angka: ")
	fmt.Scanln(&angka)
	
	for i := 1; i <= angka; i++ {
		for j := angka; j >= 1; j-- {
			if i * j == angka {
				fmt.Println(i)
			}
		}
	}
}