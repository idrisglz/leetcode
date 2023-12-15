package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l0 := &ListNode{Val: 1, Next: nil}
	l1 := &ListNode{Val: 5, Next: l0}
	result := addTwoNumbers(l1, l1)
	fmt.Println(result.Val, result.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return calculateNextNode(l1, l2, 0)
}

func calculateNextNode(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	sum := carry

	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}

	carry = sum / 10
	sum %= 10

	return &ListNode{Val: sum, Next: calculateNextNode(l1, l2, carry)}
}
