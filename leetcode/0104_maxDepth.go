package main

import "fmt"

// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=leetcode-75

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func test_maxDepth() {
	node5 := TreeNode{Val: 5}
	node4 := TreeNode{Val: 4}
	node3 := TreeNode{Val: 3}
	node2 := TreeNode{Val: 2}
	node1 := TreeNode{Val: 1}

	node3.Left = &node1
	node3.Right = &node5

	node5.Left = &node2
	node5.Right = &node4

	fmt.Println(maxDepth2(&node3), " ")
}
