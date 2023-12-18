package main

import "fmt"

var cache = map[int]int{}

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if val, ok := cache[n]; ok {
		return val
	}

	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return cache[n]
}
