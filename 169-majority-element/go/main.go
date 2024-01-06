package main

import "fmt"

func main() {
	result := majorityElement([]int{2, 3, 3})
	fmt.Println(result)
}

func majorityElement(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++

		if counts[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}
