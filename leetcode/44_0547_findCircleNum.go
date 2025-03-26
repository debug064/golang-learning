package main

import (
	"fmt"
	"strings"
	"time"
)

// Graphs - DFS
// https://leetcode.com/problems/number-of-provinces/description/?envType=study-plan-v2&envId=leetcode-75

func findCircleNum(isConnected [][]int) int {
	res := 0
	n := len(isConnected)
	var mark func(i int, j int)
	mark = func(i int, j int) {
		if isConnected[i][j] == 2 {
			return
		}
		isConnected[i][j] = 2
		j++
		for ; j < n; j++ {
			if isConnected[i][j] == 1 {
				isConnected[i][j] = 2
				if i != j {
					mark(j, 0)
				}
			} else {
				isConnected[i][j] = 2
			}
		}
	}

	for i := 0; i < n; i++ {
		if isConnected[i][0] != 2 {
			mark(i, 0)
			res++
		}
	}
	return res
}

func test_findCircleNum() {
	testCases := []struct {
		isConnected [][]int
		expected    int
	}{
		{isConnected: [][]int{
			{1, 0, 0, 1},
			{0, 1, 1, 0},
			{0, 1, 1, 1},
			{1, 0, 1, 1}}, expected: 1},
		{isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, expected: 2},
		{isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, expected: 3},
	}

	for _, tc := range testCases {
		//fmt.Printf("[%v] \n", tc.n)
		start := time.Now()
		res := findCircleNum(tc.isConnected)
		fmt.Printf("It took %v\n", time.Since(start))
		fmt.Printf("expected [%v] = %v\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
