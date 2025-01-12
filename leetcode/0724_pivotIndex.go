package main

import "fmt"

// Prefix Sum
// https://leetcode.com/problems/find-pivot-index/description/?envType=study-plan-v2&envId=leetcode-75

func pivotIndex(nums []int) int {
	sumLeft := 0
	sumRight := 0
	i := 1
	for ; i < len(nums); i++ {
		sumRight += nums[i]
	}
	i = 1
	for ; i < len(nums); i++ {
		if sumLeft == sumRight {
			return i - 1
		}
		sumLeft += nums[i-1]
		sumRight -= nums[i]
	}
	if sumLeft == sumRight {
		return i - 1
	}
	return -1
}

func testPivotIndex() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{-1, -1, 0, 1, 1, 0}, expected: 5},
		{nums: []int{1, 7, 3, 6, 5, 6}, expected: 3},
		{nums: []int{1, 2, 3}, expected: -1},
		{nums: []int{2, 1, -1}, expected: 0},
	}

	for _, tc := range testCases {
		fmt.Printf("%v \n", tc.nums)
		res := pivotIndex(tc.nums)
		fmt.Printf("expected %d = %d\n", tc.expected, res)
	}
}
