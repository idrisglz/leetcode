package main

import "fmt"

func main() {
	fmt.Println("121:", isPalindrome(121))
	fmt.Println("-121:", isPalindrome(-121))
	fmt.Println("0:", isPalindrome(0))
	fmt.Println("1230:", isPalindrome(1230))
	fmt.Println("1221:", isPalindrome(1221))
	fmt.Println("9:", isPalindrome(9))
}

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	if x < 10 {
		return true
	}

	var y int
	for x > y {
		y = y*10 + x%10
		x /= 10
	}
	return x == y || x == y/10
}
