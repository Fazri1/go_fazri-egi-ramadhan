package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	// your code here
	result := ""
	if len(a) < len(b) && strings.Contains(b, a) || len(a) == len(b) {
		result = a
	} else if len(a) > len(b) && strings.Contains(a, b) {
		result = b
	} else {
		result = "Tidak mengandung substring"
	}
	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}