// Prioritas 1-3
package main

import "fmt"

func primeNumber(number int) bool {
	if number < 2{
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}

func getPrime(number int) int {
	prime := 0
	count := 0
	num := 2

	for count < number {
		if primeNumber(num) {
			prime = num
			count += 1
		}
		num += 1
	}
	
	return prime
}

func main() {
	fmt.Println(getPrime(1))
	fmt.Println(getPrime(5))
	fmt.Println(getPrime(8))
	fmt.Println(getPrime(9))
	fmt.Println(getPrime(10))
}