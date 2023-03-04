// Prioritas_2
package main

import (
	"fmt"
	"math"
)

func frog(jumps []int) int {
	result := make([]int, len(jumps))
	result[1] = int(math.Abs(float64(jumps[1])-float64(jumps[0])))

	for i := 2; i < len(jumps); i++ {
		total1 := result[i-1] + int(math.Abs(float64(jumps[i]) - float64(jumps[i-1])))
		total2 := result[i-2] + int(math.Abs(float64(jumps[i]) - float64(jumps[i-2])))

		if total1 < total2 {
			result[i] = total1
		} else {
			result[i] = total2
		}

	}
	
	return result[len(jumps)-1]
}

func main() {
	fmt.Println(frog([]int{10,30,40,20})) // 30
	fmt.Println(frog([]int{30,10,60,10,60,50})) // 40
}