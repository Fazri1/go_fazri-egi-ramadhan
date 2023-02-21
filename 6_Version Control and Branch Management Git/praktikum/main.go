// soal eksplorasi 
package main

import "fmt"

func romanToInt(s string) int {
    roman := map[byte]int{
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}

	prevChar := 0
	result := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentChar := roman[s[i]]

		if currentChar < prevChar {
			result -= currentChar
		} else {
			result += currentChar
		}

		prevChar = currentChar
	}

	return result
}

func main() {
	fmt.Println("Hello World!")
	number := romanToInt("Ix")

	fmt.Println(number)
}