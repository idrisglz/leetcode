package main

import "fmt"

func main() {
	result := majorityElement([]int{2, 3, 3})
	fmt.Println(result)
}

func majorityElement(nums []int) int {
	var elementCounts = map[int]int{}

	for _, num := range nums {
		elementCounts[num]++
	}

	for key, value := range elementCounts {
		if value > len(nums)/2 {
			return key
		}
	}
	return 0
}
