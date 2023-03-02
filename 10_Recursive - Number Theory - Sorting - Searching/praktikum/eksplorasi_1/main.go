package main

import "fmt"

func MaxSequence(arr []int) int {
	// your code here
	result := arr[0]
	current := result

	for i := 1; i < len(arr); i++ {
		if arr[i] > current + arr[i] {
			current = arr[i]
		} else {
			current = current+arr[i]
		}
		
		if result < current {
			result = current
		}
	}

	return result

}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))  // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))    // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))    // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))    // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))     // 12
}