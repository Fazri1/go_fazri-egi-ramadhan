// prioritas_1-2
package main

import "fmt"

func segitigaPascal(numRows int) [][]int {
	segitiga := [][]int{}
	
	for i := 0; i < numRows; i++ {
		row := make([]int, i + 1)
		
		for j := 0; j < len(row); j++ {
			if j == 0 || j == len(row)-1 {
				row[j] = 1
			} else {
				row[j] = segitiga[i-1][j-1] + segitiga[i-1][j]
			}
		}
		segitiga = append(segitiga, row)
	}
	return segitiga
}

func main() {
	fmt.Println(segitigaPascal(5))
}