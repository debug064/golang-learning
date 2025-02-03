package main

import (
	"fmt"
	"strings"
	"time"
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}

func test_lowestCommonAncestor() {

	// timer := func(name string) func() {
	// 	start := time.Now()
	// 	return func() {
	// 		fmt.Printf("%s took %v\n", name, time.Since(start))
	// 	}
	// }
	testCases := []struct {
		root1    string
		p        int
		q        int
		expected int
	}{
		{root1: "3,5,1,6,2,0,8,null,null,7,4", p: 5, q: 1, expected: 3},
		{root1: "3,5,1,6,2,0,8,null,null,7,4", p: 5, q: 4, expected: 5},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		start := time.Now()
		res := lowestCommonAncestor(root1, FindInTree(root1, tc.p), FindInTree(root1, tc.q))
		fmt.Printf("It took %v\n", time.Since(start))
		fmt.Printf("expected [%d] = %d\n", tc.expected, res.Val)
	}
	fmt.Println(strings.Repeat("*", 25))
}
