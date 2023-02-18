// prioritas 2-1
// mencetak segitiga asterik

package main

import "fmt"

func main() {
	var tinggi int

	fmt.Print("Tinggi segitiga: ")
	fmt.Scanln(&tinggi)
	
	for i := 0; i < tinggi; i++ {
		for j := tinggi-2; j >= i; j-- {
			fmt.Print(" ")
			
		}
	
		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}
	
		fmt.Println()
	}
}