// Prioritas 2
// hanya hitung huruf, spesial karakter tidak dihitung
package main

import (
	"runtime"
	"sync"
	"time"
	"fmt"
)

var wg sync.WaitGroup

func frekuensiHuruf(text string, startIndex, endIndex int, ch chan map[string]int) {
	runtime.LockOSThread()
	defer wg.Done()
	
	huruf := map[string]int{}

    for _, char := range text[startIndex:endIndex] {
        if char > 65 && char < 90 || char > 97 && char < 122 {
            huruf[string(char)] += 1
        }
    }
	ch <- huruf

	time.Sleep(time.Second * 2)
	runtime.UnlockOSThread()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor, dignissim sit amet, adipiscing nec, ultricies sed, dolor."

    ch := make(chan map[string]int)
    result := make(map[string]int)
	goroutine := 5
    goroutineSize := len(text) / goroutine
	
	for i := 0; i < goroutine; i++ {
        start := i * goroutineSize
        end := start + goroutineSize
        if i == goroutine-1 {
            end = len(text)
        }

        wg.Add(1)
        go frekuensiHuruf(text, start, end, ch)
    }

	go func() {
        wg.Wait()
        close(ch)
    }()

	// gabungkan semua hasil
	for hasilGoroutine := range ch {
        for char, count := range hasilGoroutine {
            result[char] += count
        }
    }
	
	for char, count := range result {
		fmt.Println(char + ":", count)
	}
}