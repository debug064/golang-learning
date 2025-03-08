package main

import (
	"fmt"
	"strings"
)

// Linked List
// https://leetcode.com/problems/odd-even-linked-list/?envType=study-plan-v2&envId=leetcode-75
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var odd *ListNode
	var even *ListNode
	odd = head
	even = head.Next
	head2 := even
	curr := even.Next
	oe := true
	for curr != nil {
		if oe {
			odd.Next = curr
			odd = curr
		} else {
			even.Next = curr
			even = curr
		}
		oe = !oe
		curr = curr.Next
	}
	odd.Next = head2
	even.Next = nil

	return head
}

func test_oddEvenList() {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{1, 2, 3, 4, 5}, expected: []int{1, 3, 5, 2, 4}},
		{nums: []int{2, 1, 3, 5, 6, 4, 7}, expected: []int{2, 3, 6, 7, 1, 5, 4}},
	}

	makeList := func(nums []int) *ListNode {
		var list, node *ListNode
		for _, val := range nums {
			if list == nil {
				list = &ListNode{Val: val}
				node = list
				continue
			}
			node.Next = &ListNode{Val: val}
			node = node.Next
		}
		return list
	}

	for _, tc := range testCases {
		fmt.Printf("%v \n", tc.nums)
		list := makeList(tc.nums)
		res := oddEvenList(list)
		fmt.Printf("expected %v = ", tc.expected)
		print_list(res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
