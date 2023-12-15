package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var result int

	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var prev rune

	for _, curr := range s {
		result += values[curr]
		if values[prev] < values[curr] {
			result -= (values[prev] * 2)
		}
		prev = curr
	}
	return result
}
