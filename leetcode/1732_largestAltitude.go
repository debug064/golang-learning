package main

import "fmt"

// Prefix Sum
// https://leetcode.com/problems/find-the-highest-altitude/description/?envType=study-plan-v2&envId=leetcode-75

func largestAltitude(gain []int) int {
	res := 0
	val := 0
	for _, n := range gain {
		val += n
		if res < val {
			res = val
		}
	}
	return res
}

func testLargestAltitude() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{-5, 1, 5, 0, -7}, expected: 1},
		{nums: []int{-4, -3, -2, -1, 4, 3, 2}, expected: 0},
	}

	for _, tc := range testCases {
		fmt.Printf("%v \n", tc.nums)
		res := largestAltitude(tc.nums)
		fmt.Printf("expected %d = %d\n", tc.expected, res)
	}
}
