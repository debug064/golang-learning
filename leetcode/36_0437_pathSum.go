package main

import (
	"fmt"
	"strings"
)

// Binary Tree - DFS
// https://leetcode.com/problems/path-sum-iii/description/?envType=study-plan-v2&envId=leetcode-75

/**
 * Definition for a binary tree node.
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	// absInt := func(n int64) int64 {
	// 	if n < 0 {
	// 		return -n
	// 	}
	// 	return n
	// }
	var dfs func(root *TreeNode, targetSum int64, nums []int)
	dfs = func(root *TreeNode, targetSum int64, nums []int) {
		sum := int64(0)
		nums = append(nums, root.Val)
		for i := len(nums) - 1; i >= 0; i-- {
			sum += int64(nums[i])
			if sum == targetSum {
				count++
			}
		}

		if root.Left != nil {
			dfs(root.Left, targetSum, nums)
		}
		if root.Right != nil {
			dfs(root.Right, targetSum, nums)
		}
	}

	if root != nil {
		dfs(root, int64(targetSum), []int{})
	}
	return count
}

// looks like O^2
func pathSum_2(root *TreeNode, targetSum int) int {
	res := 0
	var checkSum func(node *TreeNode, pathSum int)
	checkSum = func(node *TreeNode, pathSum int) {
		if node == nil {
			return
		}
		pathSum += node.Val
		if pathSum == targetSum {
			res++
		}
		checkSum(node.Left, pathSum)
		checkSum(node.Right, pathSum)
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		checkSum(node, 0)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func test_pathSum() {

	testCases := []struct {
		root1    string
		target   int
		expected int
	}{
		{root1: "1,-2,-3", target: -2, expected: 2},
		{root1: "-2,null,-3", target: -5, expected: 1},
		{root1: "1", target: 0, expected: 0},
		{root1: "10,5,-3,3,2,null,11,3,-2,null,1", target: 8, expected: 3},
		{root1: "5,4,8,11,null,13,4,7,2,null,null,5,1", target: 22, expected: 3},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		res := pathSum(root1, tc.target)
		fmt.Printf("expected [%d] = %d\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		res := pathSum_2(root1, tc.target)
		fmt.Printf("expected [%d] = %d\n", tc.expected, res)
	}
}
