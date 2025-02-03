package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func stringToTree(str string) *TreeNode {
	root := &TreeNode{}
	tokens := strings.Split(str, ",")
	if len(tokens) == 0 || tokens[0] == "null" {
		return root
	}
	n, _ := strconv.Atoi(tokens[0])
	root.Val = n
	leafs := []*TreeNode{}
	leafs2 := []*TreeNode{}
	leafs = append(leafs, root)
	for i := 1; i < len(tokens); i++ {
		if tokens[i] != "null" {
			//leafs = leafs[1:]
			n, _ = strconv.Atoi(tokens[i])
			leafs[0].Left = &TreeNode{Val: n}
			leafs2 = append(leafs2, leafs[0].Left)
		}
		i++
		if i >= len(tokens) {
			break
		}
		if tokens[i] != "null" {
			//leafs = leafs[1:]
			n, _ = strconv.Atoi(tokens[i])
			leafs[0].Right = &TreeNode{Val: n}
			leafs2 = append(leafs2, leafs[0].Right)
		}
		leafs = leafs[1:]
		if len(leafs) == 0 {
			leafs, leafs2 = leafs2, leafs
		}
	}
	return root
}

func FindInTree(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}
	res := FindInTree(node.Left, val)
	if res == nil {
		res = FindInTree(node.Right, val)
	}
	return res
}
