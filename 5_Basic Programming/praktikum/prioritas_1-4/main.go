// Prioritas 1-4
// mencetak angka dari 1 sampai 100 dan untuk kelipatan '3' cetak "Fizz" sebagai ganti angka, dan untuk kelipatan '5' cetak "Buzz”.

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i % 3 == 0 {
			fmt.Print("Fizz ")
		} else if i % 5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
}