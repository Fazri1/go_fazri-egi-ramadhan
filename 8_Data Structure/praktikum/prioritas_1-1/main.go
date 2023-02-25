// Prioritas 1-1
package main

import (
	"fmt"
)

func arrayMerge(arrayA, arrayB []string) []string {
	arrayMerged := append(arrayA, arrayB...)
    values := map[string]bool{}
	result := []string{}

	for _, v := range arrayMerged {
		if !values[v] {
			values[v] = true
			result = append(result, v)
		}
	}

	return result
}

func main() {
	fmt.Println(arrayMerge([]string{"king","devil jin","akuma"}, []string{"eddie","steve","geese"}))
	fmt.Println(arrayMerge([]string{"sergei","jin"}, []string{"jin","steve","bryan"}))
	fmt.Println(arrayMerge([]string{"alisa","yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(arrayMerge([]string{}, []string{"devil jin","sergei"}))
	fmt.Println(arrayMerge([]string{"hwoarang"},[]string{}))
	fmt.Println(arrayMerge([]string{},[]string{}))
}