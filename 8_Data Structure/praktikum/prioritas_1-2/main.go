// Prioritas 1-2
package main

import "fmt"

func mapping(slice []string) map[string]int {
	count := map[string]int{}

	for _, v := range slice {
		count[v] += 1
	}

	return count
}

func main() {
	fmt.Println(mapping([]string{"asd","qwe","asd","adi","qwe","qwe"}))
	fmt.Println(mapping([]string{"asd","qwe","asd"}))
	fmt.Println(mapping([]string{}))
}