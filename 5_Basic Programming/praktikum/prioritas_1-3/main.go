// Prioritas 1-3
// menentukan grade dari sebuah nilai

package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Nilai: ")
	fmt.Scan(&nilai)

	switch {
	case nilai >= 80 && nilai <= 100:
		fmt.Println("A")
	case nilai >= 65 && nilai <= 79:
		fmt.Println("B")
	case nilai >= 50 && nilai <= 64:
		fmt.Println("C")
	case nilai >= 35 && nilai <= 49:
		fmt.Println("D")
	case nilai >= 0 && nilai <= 34:
		fmt.Println("E")
	default:
		fmt.Println("Nilai Invalid")
	}
}