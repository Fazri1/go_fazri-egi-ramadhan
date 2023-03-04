package main

import "fmt" 

func mergeArr(left, right[]int) []int{
	var i, j int
	result := []int{}

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}

func mergeSort(arr []int) []int{
	if len(arr) < 2 {
		return arr
	}
	var mid int
	if len(arr) % 2 == 0 {
		mid = len(arr) / 2
	} else {
		mid = (len(arr) + 1) / 2
	}
	mergeLeft := mergeSort(arr[:mid])
	mergeRight := mergeSort(arr[mid:])
	return mergeArr(mergeLeft, mergeRight)
	
}

func MaximumBuyProduct(money int, productPrice []int) {
	// your code here
	sortedProduct := mergeSort(productPrice)
	productBuy := 0

	for i := range sortedProduct {
		if money >= sortedProduct[i] {
			productBuy += 1
			money -= sortedProduct[i]
		} else {
			break
		}
	}

	fmt.Println(productBuy)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // 0

}