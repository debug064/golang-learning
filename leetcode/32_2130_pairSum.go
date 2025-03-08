package main

import (
	"fmt"
	"strings"
)

// Linked List
// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/?envType=study-plan-v2&envId=leetcode-75
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func pairSum(head *ListNode) int {
	var slow *ListNode
	var fast *ListNode
	slow = head
	fast = head.Next
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head2 := slow.Next
	slow = slow.Next.Next
	head2.Next = nil
	for slow != nil {
		tmp := head2
		head2 = slow
		slow = slow.Next
		head2.Next = tmp
	}
	res := head2.Val + head.Val
	head = head.Next
	head2 = head2.Next
	for head2 != nil {
		tmp := head2.Val + head.Val
		if tmp > res {
			res = tmp
		}
		head = head.Next
		head2 = head2.Next
	}

	return res
}

func test_pairSum() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{5, 4, 2, 1}, expected: 6},
		{nums: []int{4, 2, 2, 3}, expected: 7},
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
		res := pairSum(list)
		fmt.Printf("expected %v = ", tc.expected)
		fmt.Println(res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
