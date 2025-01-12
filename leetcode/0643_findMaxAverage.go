package main

import (
	"fmt"
	"math"
	"strings"
)

// Sliding Window
// https://leetcode.com/problems/maximum-average-subarray-i/?envType=study-plan-v2&envId=leetcode-75

func findMaxAverage(nums []int, k int) float64 {
	if k <= 0 {
		return 0
	}
	res := math.MinInt
	var sum int = 0
	w := 0
	for i, val := range nums {
		sum += val
		w++
		if w == k {
			if res < sum {
				res = sum
			}
			sum -= nums[i-k+1]
			w--
		}
	}
	return float64(res) / float64(k)
}

func testFindMaxAverage() {
	testCases := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{nums: []int{-1}, k: 1, expected: -1.0},
		{nums: []int{1, 12, -5, -6, 50, 3}, k: 4, expected: 12.75},
		{nums: []int{5}, k: 1, expected: 5.0},
	}

	for _, val := range testCases {
		fmt.Println(val.nums)
		res := findMaxAverage(val.nums, val.k)
		fmt.Printf("expected %f =  %f\n", val.expected, res)
		fmt.Println(strings.Repeat("*", 25))
	}
}
