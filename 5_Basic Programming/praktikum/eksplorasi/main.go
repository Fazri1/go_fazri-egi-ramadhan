// eksplorasi
// memeriksa apakah sebuah kata adalah palindrome

package main

import (
	"fmt"	
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("Apakah Palindrome?")
	fmt.Print("Masukkan kata: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kata := strings.TrimSpace(scanner.Text())

	fmt.Println("captured:", kata)

	pali := ""
	
	for _, v := range kata {
		pali = string(v) + pali
	}
	
	if pali == kata {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}