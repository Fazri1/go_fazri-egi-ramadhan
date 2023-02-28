// Prioritas 2
package main

import (
	"fmt"
)
func caesar(offset int, input string) string {
	// your code here
	result := ""
	ascii := []rune(input)
	newOffset := offset

	if newOffset > 26 {
		newOffset = offset % 26
	}

	for i := range ascii {
		if ascii[i] < 97 {
			result += string(ascii[i])
		} else if ascii[i] + rune(newOffset) > 122 {
			offset = newOffset - (122 - int(ascii[i]))
			newAscii := 96
			result += string(newAscii + offset)
		} else {
			result += string(ascii[i] + rune(newOffset))
		}
	}
	
	return result
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // def
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi 
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza 
  	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl 
}