package main

import (
	"fmt"
	"strings"
	"time"
)

// Binary Search Tree
// https://leetcode.com/problems/search-in-a-binary-search-tree/description/?envType=study-plan-v2&envId=leetcode-75

/**
 * Definition for a binary tree node.
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

func test_searchBST() {

	// timer := func(name string) func() {
	// 	start := time.Now()
	// 	return func() {
	// 		fmt.Printf("%s took %v\n", name, time.Since(start))
	// 	}
	// }
	testCases := []struct {
		root1    string
		val      int
		expected string
	}{
		{root1: "4,2,7,1,3", val: 2, expected: "2,1,3"},
		{root1: "4,2,7,1,3", val: 5, expected: ""},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		start := time.Now()
		res := searchBST(root1, tc.val)
		fmt.Printf("It took %v\n", time.Since(start))
		fmt.Printf("expected [%v] = %v\n", tc.expected, treeToString(res))
	}
	fmt.Println(strings.Repeat("*", 25))
}
