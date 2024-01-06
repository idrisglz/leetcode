package main

func main() {

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
