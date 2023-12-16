package main

import (
	"testing"
)

type TestCase struct {
	Input    []int
	Expected int
}

func TestTrap(t *testing.T) {
	testCases := []TestCase{
		{Input: []int{0, 2, 1, 1}, Expected: 0},
		{Input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, Expected: 6},
		{Input: []int{5, 4, 1, 2}, Expected: 1},
		{Input: []int{3, 2, 1, 2, 1}, Expected: 1},
		{Input: []int{2, 1, 1, 3}, Expected: 2},
		{Input: []int{4, 3, 3, 9, 3, 0, 9, 2, 8, 3}, Expected: 23},
		{Input: []int{2, 1, 2, 1, 3}, Expected: 2},
		{Input: []int{3, 1, 2, 1, 3}, Expected: 5},
		{Input: []int{3, 2, 4, 5}, Expected: 1},
		{Input: []int{4, 2, 0, 3, 2, 5}, Expected: 9},
		{Input: []int{4, 2, 3}, Expected: 1},
		{Input: []int{0, 7, 1, 4, 6}, Expected: 7},
	}

	for _, tc := range testCases {
		if actual := trap(tc.Input); actual != tc.Expected {
			t.Errorf("Test case failed %v - expected %d, got %d", tc.Input, tc.Expected, actual)
		}
	}

}
