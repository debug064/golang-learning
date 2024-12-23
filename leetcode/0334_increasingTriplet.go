package main

import "fmt"

// https://leetcode.com/problems/increasing-triplet-subsequence/description/?envType=study-plan-v2&envId=leetcode-75

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	min := nums[0]
	mid := 1 << 31
	for _, num := range nums {
		fmt.Printf("%d - ", num)
		if num <= min {
			min = num
		} else if num <= mid {
			mid = num
		} else {
			return true
		}

		fmt.Printf("%d; %d\n", min, mid)
	}
	return false
}

func testIncreasingTriplet() {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{nums: []int{1, 2, 3, 4, 5}, expected: true},
		{nums: []int{5, 4, 3, 2, 1}, expected: false},
		{nums: []int{2, 1, 5, 0, 4, 6}, expected: true},
		{nums: []int{1, 5, 0, 4, 1, 3}, expected: true},
		{nums: []int{20, 100, 10, 12, 5, 13}, expected: true},
		{nums: []int{0, 4, 2, 1, 0, -1, -3}, expected: false},
	}

	for _, tc := range testCases {
		fmt.Printf("%v\n", tc.nums)
		res := increasingTriplet(tc.nums)
		fmt.Printf("expected %t = %t\n", tc.expected, res)
	}
}
