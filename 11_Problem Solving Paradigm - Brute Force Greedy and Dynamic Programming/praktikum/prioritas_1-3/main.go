// prioritas_1-3
package main

import "fmt"

func save(n int, saved map[int]int) int {
	if v, exist := saved[n]; exist {
		return v
	}

	if n <= 1 {
		saved[n] = n
		return n 
	}
	
	saved[n] = save(n-1, saved) + save(n-2, saved)
	return saved[n]
}

func fibo(n int) int {
	saved := map[int]int{}

	return save(n, saved)
	
}

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
}