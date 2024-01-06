package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumberXOR([]int{2, 2, 1}))

}

func singleNumber(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	for key, val := range counts {
		if val == 1 {
			return key
		}
	}
	return -1
}

// Using XOR
func singleNumberXOR(nums []int) int {
	var result int

	for _, num := range nums {
		result ^= num
	}

	return result
}
