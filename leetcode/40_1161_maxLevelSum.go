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
func maxLevelSum(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	sum := 0
	maxSum := root.Val
	res = 1
	j := 1
	var nodes []*TreeNode
	var nodes2 []*TreeNode
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		sum += nodes[0].Val
		if nodes[0].Left != nil {
			nodes2 = append(nodes2, nodes[0].Left)
		}
		if nodes[0].Right != nil {
			nodes2 = append(nodes2, nodes[0].Right)
		}
		nodes = nodes[1:]
		if len(nodes) == 0 {
			nodes, nodes2 = nodes2, nodes
			if sum > maxSum {
				maxSum = sum
				res = j
			}
			sum = 0
			j++
		}
	}
	return res
}

func test_maxLevelSum() {

	// timer := func(name string) func() {
	// 	start := time.Now()
	// 	return func() {
	// 		fmt.Printf("%s took %v\n", name, time.Since(start))
	// 	}
	// }
	testCases := []struct {
		root1    string
		expected int
	}{
		{root1: "1,7,0,7,-8,null,null", expected: 2},
		{root1: "989,null,10250,98693,-89388,null,null,null,-32127", expected: 2},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		start := time.Now()
		res := maxLevelSum(root1)
		fmt.Printf("It took %v\n", time.Since(start))
		fmt.Printf("expected [%v] = %v\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
