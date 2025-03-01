package main

import (
	"fmt"
	"strings"
)

// Linked List
// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/description/?envType=study-plan-v2&envId=leetcode-75
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var node *ListNode
	var mid **ListNode
	mid = &head
	node = head
	move := false
	for node != nil {
		if move {
			mid = &(*mid).Next
		}
		move = !move
		node = node.Next
	}
	if (*mid).Next != nil {
		(*mid).Val = (*mid).Next.Val
		(*mid).Next = (*mid).Next.Next
	} else {
		*mid = nil
	}
	return head
}

func deleteMiddle2(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	var slow, fast *ListNode
	slow = head
	fast = head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func test_deleteMiddle() {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{2}, expected: []int{}},
		{nums: []int{1, 3, 4, 7, 1, 2, 6}, expected: []int{1, 3, 4, 1, 2, 6}},
		{nums: []int{1, 2, 3, 4}, expected: []int{1, 2, 4}},
		{nums: []int{2, 1}, expected: []int{2}},
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
		res := deleteMiddle(list)
		fmt.Printf("expected %v = ", tc.expected)
		print_list(res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, tc := range testCases {
		fmt.Printf("%v \n", tc.nums)
		list := makeList(tc.nums)
		res := deleteMiddle2(list)
		fmt.Printf("expected %v = ", tc.expected)
		print_list(res)
	}
}

func print_list(head *ListNode) {
	var node1 *ListNode
	node1 = head
	for node1 != nil {
		fmt.Print(node1.Val, " ")
		node1 = node1.Next
	}
	fmt.Println()
}
