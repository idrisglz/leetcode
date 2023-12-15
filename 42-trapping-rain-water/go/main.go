package main

import "fmt"

func main() {
	num := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println(num)
}

func trap(height []int) int {
	var sum int
	for i := 0; i < len(height); {
		holder, end := canHold(height, i)
		if holder {
			sum += calculateWater(height, i, end)
			i = end
		} else {
			i++
		}
	}
	return sum
}

func calculateWater(height []int, currIndex int, endIndex int) int {
	var maxHeight int
	if height[currIndex] < height[endIndex] {
		maxHeight = height[currIndex]
	} else {
		maxHeight = height[endIndex]
	}
	length := endIndex - currIndex - 1

	var occupied int

	for i := currIndex + 1; i < endIndex; i++ {
		occupied += height[i]
	}
	capacity := length*maxHeight - occupied
	return capacity
}

func canHold(height []int, currIndex int) (bool, int) {
	if currIndex == len(height)-1 {
		return false, currIndex + 1
	}

	if height[currIndex+1] >= height[currIndex] {
		return false, currIndex + 1
	}

	var currentMaxEnd int
	var endIndex int
	for i := currIndex + 2; i < len(height); i++ {
		if height[i] >= height[currIndex] {
			return true, i
		}

		if height[i] > height[currIndex+1] && height[i] > currentMaxEnd {
			currentMaxEnd = height[i]
			endIndex = i
		}
	}
	return endIndex > currIndex+1, endIndex
}
