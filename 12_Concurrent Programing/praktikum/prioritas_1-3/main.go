// Prioritas 1-3
package main

import (
	"fmt"
)

func kelipatanTiga(num int, ch chan int) {
    for i := 1; i <= num; i++ {
        if i%3 == 0 {
            ch <- i
        }
    }
	close(ch)
}

func main() {
    ch := make(chan int, 3)

    go kelipatanTiga(100, ch)

    for n := range ch {
        fmt.Println(n)
    }
	
}