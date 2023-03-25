// Prioritas 1-2
package main

import (
	"fmt"
	"time"
)

func main(){
	number := 100
	ch := make(chan int)

	go func(max int) {
		for i := 1; i <= max; i++ {
			if i % 3 == 0 {
				ch <- i
			}	
		}
	}(number)

	go func(n int) {
		for i := 1; i <= n; i++ {
			num := <- ch
			fmt.Println("Data dari channel:", num)
		}
	}(number)
	<-time.After(time.Second *3)
}