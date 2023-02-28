// Prioritas 1-4
package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min, max int) {
	min = *numbers[0]
	max = *numbers[0]
	
	for i := range numbers {
		if min > *numbers[i] {
			min = *numbers[i]
		} else if max < *numbers[i] {
			max = *numbers[i]
		}
	}
	
	return min, max
}

func main(){
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	
	fmt.Println()
	fmt.Println("Nilai min:", min)
	fmt.Println("Nilai max:", max)
}