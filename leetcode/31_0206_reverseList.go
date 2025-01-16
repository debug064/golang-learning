package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var node, next *ListNode
	for head != nil {
		node = head
		head = head.Next
		node.Next = next
		next = node
	}
	return node
}

func test_reverseList() {
	//var node1 ListNode
	node5 := ListNode{Val: 5}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	print_reverseList(&node1)
	print_reverseList(reverseList(&node1))
}

func print_reverseList(head *ListNode) {
	var node1 *ListNode
	node1 = head
	for node1 != nil {
		fmt.Print(node1.Val, " ")
		node1 = node1.Next
	}
	fmt.Println()
}
