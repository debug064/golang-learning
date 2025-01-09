package main

import "fmt"

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/?envType=study-plan-v2&envId=leetcode-75

func longestSubarray(nums []int) int {
	res := 0
	l := 0
	r := 0
	val := 0
	for i, n := range nums {
		if n == 1 {
			val++
			if val > res {
				res = val
			}
			continue
		}
		for ; l < r; l++ {
			if nums[l] == 1 {
				val--
			}
		}
		r = i

	}
	return res - nums[r]
}

func testLongestSubarray() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{1, 0, 0, 0, 0}, expected: 1},
		{nums: []int{1, 1, 0, 1}, expected: 3},
		{nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, expected: 5},
		{nums: []int{1, 1, 1}, expected: 2},
	}

	for _, tc := range testCases {
		fmt.Printf("%v \n", tc.nums)
		res := longestSubarray(tc.nums)
		fmt.Printf("expected %d = %d\n", tc.expected, res)
	}
}
