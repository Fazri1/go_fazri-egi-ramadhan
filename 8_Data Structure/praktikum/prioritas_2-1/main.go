// Prioritas 2
package main

import "fmt"

func pairSum(arr []int, target int) []int {
	result := []int{}

	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] + arr[j] == target && len(result) <= 2 {
				result = append(result, i,j)
			}
		}
	}

	return result
}

func main() {
	fmt.Println(pairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(pairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(pairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(pairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(pairSum([]int{1, 5, 6, 7}, 6))
}