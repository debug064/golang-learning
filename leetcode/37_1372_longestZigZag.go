package main

import (
	"fmt"
	"strings"
)

// Binary Tree - DFS
// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/?envType=study-plan-v2&envId=leetcode-75

/**
 * Definition for a binary tree node.
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func longestZigZag(root *TreeNode) int {
	max_count := 0
	//var dfs func(root *TreeNode)
	var dfs_count func(root *TreeNode, count int, dir bool)
	dfs_count = func(root *TreeNode, count int, dir bool) {
		if count > max_count {
			max_count = count
		}
		if !dir {
			if root.Left != nil {
				dfs_count(root.Left, count+1, !dir)
			}
			if root.Right != nil {
				dfs_count(root.Right, 0, !dir)
			}
		} else {
			if root.Right != nil {
				dfs_count(root.Right, count+1, !dir)
			}
			if root.Left != nil {
				dfs_count(root.Left, 0, !dir)
			}
		}
	}
	dfs_count(root, 0, true)
	dfs_count(root, 0, false)

	return max_count
}

func test_longestZigZag() {

	testCases := []struct {
		root1    string
		expected int
	}{
		{root1: "1,null,1,1,1,null,null,1,1,null,1,null,null,null,1", expected: 3},
		{root1: "1,1,1,null,1,null,null,1,1,null,1", expected: 4},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		res := longestZigZag(root1)
		fmt.Printf("expected [%d] = %d\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
