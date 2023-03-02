// Prioritas 2
// tambahan : mengeluarkan kartu yang jumlahnya lebih besar

package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	card := [][]int{}
	for i := range cards {
		if cards[i][0] == deck[0] || cards[i][0] == deck[1] || cards[i][1] == deck[0] || cards[i][1] == deck[1] {
			card = append(card, cards[i])
		} 
	}
	
	if len(card) >= 1 {
		bigCard := []int{card[0][0], card[0][1]}
		for i := 0; i < len(card)-1; i++ {	
			if bigCard[0] + bigCard[1] < card[i+1][0] + card[i+1][1] {
				bigCard = card[i+1]
			}
		}
		return bigCard
	}
	return "Tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}