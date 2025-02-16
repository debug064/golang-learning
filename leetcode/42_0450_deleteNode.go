package main

import (
	"fmt"
	"strings"
	"time"
)

// Binary Search Tree
// https://leetcode.com/problems/delete-node-in-a-bst/?envType=study-plan-v2&envId=leetcode-75

/**
 * Definition for a binary tree node.
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func deleteNode2(root *TreeNode, key int) *TreeNode {
	var searchBST func(root **TreeNode, val int) **TreeNode
	searchBST = func(root **TreeNode, val int) **TreeNode {
		if (*root) == nil {
			return nil
		}
		if val == (*root).Val {
			return root
		}
		if val < (*root).Val {
			return searchBST(&(*root).Left, val)
		} else {
			return searchBST(&(*root).Right, val)
		}
	}
	res := searchBST(&root, key)
	if res == nil {
		return root
	}
	if (*res).Right == nil && (*res).Left == nil {
		*res = nil
		return root
	}
	if (*res).Right == nil && (*res).Left != nil {
		*res = (*res).Left
		return root
	}
	if (*res).Right != nil && (*res).Left == nil {
		*res = (*res).Right
		return root
	}

	if (*res).Right.Left == nil {
		(*res).Val = (*res).Right.Val
		(*res).Right = (*res).Right.Right
		return root
	}

	tmp := (*res).Right.Left
	tmpP := &(*res).Right.Left
	for tmp.Left != nil {
		tmpP = &tmp.Left
		tmp = tmp.Left

	}
	(*res).Val = tmp.Val
	*tmpP = tmp.Right
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			tmp := root.Right
			for tmp.Left != nil {
				tmp = tmp.Left
			}
			root.Val = tmp.Val
			root.Right = deleteNode(root.Right, tmp.Val)
		}
	}
	return root
}

func test_deleteNode() {

	// timer := func(name string) func() {
	// 	start := time.Now()
	// 	return func() {
	// 		fmt.Printf("%s took %v\n", name, time.Since(start))
	// 	}
	// }
	testCases := []struct {
		root1    string
		key      int
		expected string
	}{
		{root1: "3,2,5,null,null,4,10,null,null,8,15,7", key: 5, expected: "3,2,7,null,null,4,10,null,null,8,15"},
		{root1: "5,3,6,2,4,null,7", key: 3, expected: "5,4,6,2,null,null,7"},
		{root1: "50,30,70,null,40,60,80", key: 50, expected: "60,30,70,null,40,null,80"},
		{root1: "0", key: 0, expected: ""},
		{root1: "5,3,6,2,4,null,7", key: 7, expected: "5,3,6,2,4"},
		{root1: "5,3,6,2,4,null,7", key: 0, expected: "5,3,6,2,4,null,7"},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		start := time.Now()
		res := deleteNode(root1, tc.key)
		fmt.Printf("It took %v\n", time.Since(start))
		fmt.Printf("expected [%v] = %v\n", tc.expected, treeToString(res))
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		root1 := stringToTree(tc.root1)
		start := time.Now()
		res := deleteNode2(root1, tc.key)
		fmt.Printf("It took %v\n", time.Since(start))
		fmt.Printf("expected [%v] = %v\n", tc.expected, treeToString(res))
	}
}
