// Prioritas 1-3
package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	unique := map[int]int{}
	result := []int{}

	for _,v := range angka {
		value,_ := strconv.Atoi(string(v))
		unique[value] += 1
	}

	for k,v := range unique {
		if v == 1 {
			result = append(result, k)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}