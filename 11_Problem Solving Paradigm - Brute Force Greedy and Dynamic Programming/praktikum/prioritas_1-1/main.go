// prioritas_1-1
package main

import (
	"fmt"
	"strconv"
)

func intToBinary(number int) string {
	result := ""
	if number == 0 {
		return "0"
	}

	for number > 0 {
		bin := number % 2
		result = strconv.Itoa(bin) + result
		number /= 2
	}
	return result
}

func binaryList(n int) []string {
	ans := make([]string, n+1)

	for i := 0; i <= n; i++ {
		binary := intToBinary(i)
		num := ""
		
		for _, v := range binary {
			if v == '1' {
				num += "1"
			} else {
				num += "0"
			}
		}

		ans[i] = num
	}

	return ans
}

func main() {
	fmt.Println(binaryList(2))
	fmt.Println(binaryList(3))
}