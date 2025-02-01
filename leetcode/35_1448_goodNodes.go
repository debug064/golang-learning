package main

import (
	"fmt"
	"math"
	"strings"
)

// Binary Tree - DFS
// https://leetcode.com/problems/count-good-nodes-in-binary-tree/description/?envType=study-plan-v2&envId=leetcode-75

/**
 * Definition for a binary tree node.
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func goodNodes(root *TreeNode) int {
	count := 0
	var dfs func(root *TreeNode, maxNum int)
	dfs = func(root *TreeNode, maxNum int) {
		if root.Val >= maxNum {
			maxNum = root.Val
			count++
		}
		if root.Left != nil {
			dfs(root.Left, maxNum)
		}
		if root.Right != nil {
			dfs(root.Right, maxNum)
		}
	}

	dfs(root, math.MinInt)
	return count
}

// func leafSimilar(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}
// 	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
// }

func test_goodNodes() {

	testCases := []struct {
		root1    string
		expected int
	}{
		{root1: "3,1,4,3,null,1,5", expected: 4},
		{root1: "3,3,null,4,2", expected: 3},
		{root1: "1", expected: 1},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		res := goodNodes(root1)
		fmt.Printf("expected [%d] = %d\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
