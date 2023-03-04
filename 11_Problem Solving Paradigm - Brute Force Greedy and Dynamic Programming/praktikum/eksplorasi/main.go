// Eksplorasi
package main

import (
	"fmt"
	"strconv"
)

func intToRoman(n int) string {
	result := ""
	roman := map[int][]string {
		0:[]string{"1000", "M"},
		1:[]string{"900", "CM"},
		2:[]string{"500", "D"},
		3:[]string{"400", "CD"},
		4:[]string{"100", "C"},
		5:[]string{"90", "XC"},
		6:[]string{"50", "L"},
		7:[]string{"40", "XL"},
		8:[]string{"10", "X"},
		9:[]string{"9", "IX"},
		10:[]string{"5", "V"},
		11:[]string{"4", "IV"},
		12:[]string{"1", "I"},
		
	}

	for i := 0; i < len(roman); i++ {
		num,_ := strconv.Atoi(roman[i][0])
		for n >= num {
			result += roman[i][1]
			n -= num
		}
	}

	return result
}

func main() {
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(23))
	fmt.Println(intToRoman(2021))
	fmt.Println(intToRoman(1646))
}