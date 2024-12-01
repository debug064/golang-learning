package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	max_sum := math.MinInt32
	for _, val := range nums {
		sum += val
		max_sum = max(max_sum, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return max_sum
}

type TestData53 struct {
	nums     []int
	expected int
}

func test_maxSubArray() {
	var tests = []TestData53{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6},
		{nums: []int{5, 4, -1, 7, 8}, expected: 23},
		{nums: []int{-2, -1}, expected: -1},
	}

	for _, val := range tests {
		fmt.Println(val)
		res := maxSubArray(val.nums)
		fmt.Printf("expected %d = %d\n", val.expected, res)
	}
}
