package main

import (
	"fmt"
)

func main() {
	fmt.Println(trap([]int{3, 2, 4, 5}), 1)
	fmt.Println(trap([]int{2, 1, 2, 1, 3}), 2)
	fmt.Println(trap([]int{3, 1, 2, 1, 3}), 2)
	fmt.Println(trap([]int{3, 2, 1, 2, 1}), 1)
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
}

func trap(height []int) int {
	var capacity int
	start, end := 0, len(height)-1
	startMax, endMax := height[start], height[end]

	for end >= start {
		if startMax <= endMax {
			if height[start] > startMax {
				startMax = height[start]
			}
			capacity += startMax - height[start]
			start++
		} else {
			if height[end] > endMax {
				endMax = height[end]
			}
			capacity += endMax - height[end]
			end--
		}
	}
	return capacity
}
