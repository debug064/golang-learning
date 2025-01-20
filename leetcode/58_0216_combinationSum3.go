package main

import (
	"fmt"
	"slices"
	"strings"
)

// Backtracking
// https://leetcode.com/problems/combination-sum-iii/?envType=study-plan-v2&envId=leetcode-75

func sum3(res *[][]int, k int, sum int, beg int, path []int) {
	if k == 0 && sum == 0 {
		fmt.Printf("%v\n", path)
		*res = append(*res, slices.Clone(path))
		return
	}
	if k < 0 || sum < 0 {
		return
	}
	for i := beg; i <= 9; i++ {
		path = append(path, i)
		sum3(res, k-1, sum-i, i+1, path)
		path = path[:len(path)-1]
	}
}

func combinationSum3_1(k int, n int) [][]int {
	res := [][]int{}
	sum3(&res, k, n, 1, []int{})
	return res
}

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	var lsum3 func(k int, sum int, beg int, path []int)
	lsum3 = func(k int, sum int, beg int, path []int) {
		if k == 0 && sum == 0 {
			res = append(res, slices.Clone(path))
			return
		}
		if k < 0 || sum < 0 {
			return
		}
		for i := beg; i <= 9; i++ {
			path = append(path, i)
			lsum3(k-1, sum-i, i+1, path)
			path = path[:len(path)-1]
		}
	}
	lsum3(k, n, 1, []int{})
	return res
}

func testCombinationSum3() {
	testCases := []struct {
		k        int
		n        int
		expected [][]int
	}{
		{k: 3, n: 7, expected: [][]int{{1, 2, 4}}},
		{k: 3, n: 9, expected: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{k: 4, n: 1, expected: [][]int{}},
	}

	for _, tc := range testCases {
		fmt.Printf("[%d], [%d] \n", tc.k, tc.n)
		res := combinationSum3_1(tc.k, tc.n)
		fmt.Printf("expected [%v] = %v\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, tc := range testCases {
		fmt.Printf("[%d], [%d] \n", tc.k, tc.n)
		res := combinationSum3(tc.k, tc.n)
		fmt.Printf("expected [%v] = %v\n", tc.expected, res)
	}
}
