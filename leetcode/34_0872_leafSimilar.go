package main

import (
	"fmt"
	"slices"
	"strings"
)

// Binary Tree - DFS
// https://leetcode.com/problems/leaf-similar-trees/?envType=study-plan-v2&envId=leetcode-75

/**
 * Definition for a binary tree node.
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var getLeafs func(root *TreeNode) []int
	getLeafs = func(root *TreeNode) []int {
		leafs := []int{}
		if root == nil {
			return leafs
		}
		if root.Left == nil && root.Right == nil {
			return append(leafs, root.Val)
		}
		if root.Left != nil {
			leafs = append(leafs, getLeafs(root.Left)...)
		}
		if root.Right != nil {
			leafs = append(leafs, getLeafs(root.Right)...)
		}
		return leafs
	}
	leafs1 := getLeafs(root1)
	leafs2 := getLeafs(root2)
	return slices.Equal(leafs1, leafs2)
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

func test_leafSimilar() {

	testCases := []struct {
		root1    string
		root2    string
		expected bool
	}{
		{root1: "3,5,1,6,2,9,8,null,null,7,4", root2: "3,5,1,6,7,4,2,null,null,null,null,null,null,9,8", expected: true},
		{root1: "1,2,3", root2: "1,3,2", expected: false},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		root2 := stringToTree(tc.root2)
		res := leafSimilar(root1, root2)
		fmt.Printf("expected [%t] = %t\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
