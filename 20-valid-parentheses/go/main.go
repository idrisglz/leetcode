package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid(")()"))
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
			continue
		} else if len(stack) == 0 || stack[len(stack)-1] != pairs[r] {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
