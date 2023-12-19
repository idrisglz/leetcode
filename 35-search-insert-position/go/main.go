package main

func main() {
	searchInsert([]int{1, 3, 5, 6}, 5)
	searchInsertLinear([]int{1, 3, 5, 6}, 5)
}

// Binary search
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// Linear search
func searchInsertLinear(nums []int, target int) int {
	for i, val := range nums {
		if val >= target {
			return i
		}
	}
	return len(nums)
}
