// Prioritas 1-4
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	num := 5
	result := 1

	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mutex.Lock()
			result *= n
			mutex.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Println(result)
}