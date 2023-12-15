package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

// Memory efficient
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// Compute efficient
func twoSum2(nums []int, target int) []int {
	seen := make(map[int]int)
	for index, num := range nums {
		if seenIndex, ok := seen[target-num]; ok {
			return []int{seenIndex, index}
		}
		seen[num] = index
	}
	return []int{}
}
