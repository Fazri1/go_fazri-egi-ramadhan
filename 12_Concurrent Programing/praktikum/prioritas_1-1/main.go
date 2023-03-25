// Prioritas 1-1
package main

import (
	"fmt"
	"time"
)

func kelipatan(x int) {
	if x < 1 {
		fmt.Println("Masukkan bilangan positif!")
	} else {
		for i := 1; i <= 10; i++ {
			fmt.Println(x*i)
		}
	}
}

func main() {
	start := time.Now()
	fmt.Println("Start", time.Since(start))

	go kelipatan(-1)

	time.Sleep(3 * time.Second)
	fmt.Println("Stop", time.Since(start))
}