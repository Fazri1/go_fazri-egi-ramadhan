// Prioritas 1-4
package main

import "fmt"

func SimpleEquations(a,b,c int) interface{} {
	for x := 1; x <= b; x++ {
		if b%x == 0 {

			for y := 1; y <= a; y++ {
				z := a - x - y

				if x*y*z == b && (x*x) + (y*y) + (z*z) == c {
					result := []int{x, y, z}
					return result
				}
			}
		}
	}
	return "no solution"
}

func main() {
	fmt.Println(SimpleEquations(6,6,14))
	fmt.Println(SimpleEquations(1,2,3))
	
}