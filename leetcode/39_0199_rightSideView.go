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

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	var nodes []*TreeNode
	var nodes2 []*TreeNode
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		if len(nodes) == 1 {
			res = append(res, nodes[0].Val)
		}
		if nodes[0].Left != nil {
			nodes2 = append(nodes2, nodes[0].Left)
		}
		if nodes[0].Right != nil {
			nodes2 = append(nodes2, nodes[0].Right)
		}
		nodes = nodes[1:]
		if len(nodes) == 0 {
			nodes, nodes2 = nodes2, nodes
		}
	}
	return res
}

func test_rightSideView() {

	// timer := func(name string) func() {
	// 	start := time.Now()
	// 	return func() {
	// 		fmt.Printf("%s took %v\n", name, time.Since(start))
	// 	}
	// }
	testCases := []struct {
		root1    string
		expected []int
	}{
		{root1: "1,2,3,null,5,null,4", expected: []int{1, 3, 4}},
		{root1: "1,2,3,4,null,null,null,5", expected: []int{1, 3, 4, 5}},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		start := time.Now()
		res := rightSideView(root1)
		fmt.Printf("It took %v\n", time.Since(start))
		fmt.Printf("expected [%v] = %v\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
