// Eksplorasi
package main

import "fmt"

func selisihAbsolut(slice [][]int) int {
	result := 0
	diagonalL := 0
	diagonalR := 0
	matriks := true

	for i := range slice {
		if len(slice) == len(slice[i]) {
			diagonalL += slice[i][i]
		} else {
			matriks = false
		}
	}

	if matriks {
		j := 0
		for i := len(slice)-1; i >= 0; i-- {
			diagonalR += slice[j][i]
			j++
		}
		
		if diagonalL > diagonalR {
			result = diagonalL - diagonalR
			
		} else {
			result = diagonalR - diagonalL
		}
		matriks = false
	}
	return result
}

func main() {
	fmt.Println(selisihAbsolut([][]int{
		{1,2,3},
		{4,5,6},
		{9,8,9},
	}))
}