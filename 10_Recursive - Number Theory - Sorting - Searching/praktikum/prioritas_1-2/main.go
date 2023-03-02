// Prioritas 1-2
package main

import "fmt"

type pair struct {
	name string
	count int
}

func mostAppearItem(items []string) []pair {
	result := map[string]int{}
	hasil := []pair{}

	for i := range items {
		if _, exist := result[items[i]]; exist {
			result[items[i]] += 1
		} else {
			result[items[i]] = 1
		}
	}


	for k,v := range result {
		hasil = append(hasil, pair{name: k,count: v})
	}

	for i := 0; i < len(hasil); i++ {
		for j := i+1; j < len(hasil); j++ {
			if hasil[i].count > hasil[j].count {
				hasil[i], hasil[j] = hasil[j], hasil[i]
			}
		}
	}

	return hasil
}

func main() {
	// return dengan tipe data []pair

	fmt.Println(mostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(mostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(mostAppearItem([]string{"football", "basketball", "tenis"}))
}