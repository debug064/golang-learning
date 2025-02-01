package main

import (
	"fmt"
	"strings"
)

// DP - 1D
// https://leetcode.com/problems/min-cost-climbing-stairs/description/?envType=study-plan-v2&envId=leetcode-75

func rob(nums []int) int {
	res := 0
	m := make(map[int]int)
	var lrob func(i int) int
	lrob = func(i int) int {
		val, ok := m[i]
		if ok {
			return val
		}
		val = 0
		i0 := i
		i += 2
		for ; i < len(nums); i++ {
			val = max(val, lrob(i))
		}
		m[i0] = val + nums[i0]
		return val + nums[i0]
	}

	for i := 0; i < len(nums); i++ {
		res = max(res, lrob(i))
	}
	return res
}

func rob_leet(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	prev2, prev1 := 0, nums[0]
	for i := 1; i < n; i++ {
		current := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = current
	}
	return prev1
}

func testRob() {
	testCases := []struct {
		n        []int
		expected int
	}{
		{n: []int{2, 8, 1, 2}, expected: 10},
		{n: []int{1, 2, 3, 1}, expected: 4},
		{n: []int{2, 7, 9, 3, 1}, expected: 12},
		{n: []int{2, 1, 1, 2}, expected: 4},
	}

	for _, tc := range testCases {
		fmt.Printf("[%v] \n", tc.n)
		res := rob(tc.n)
		fmt.Printf("expected [%d] = %d\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, tc := range testCases {
		fmt.Printf("[%v] \n", tc.n)
		res := rob_leet(tc.n)
		fmt.Printf("expected [%d] = %d\n", tc.expected, res)
	}
}
