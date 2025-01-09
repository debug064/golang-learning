package main

import "fmt"

// https://leetcode.com/problems/container-with-most-water/?envType=study-plan-v2&envId=leetcode-75

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	aMax := 0
	for i < j {
		a := (j - i) * min(height[i], height[j])
		if a > aMax {
			aMax = a
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return aMax
}

func testMaxArea() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expected: 49},
		{nums: []int{1, 1}, expected: 1},
	}

	for _, tc := range testCases {
		fmt.Printf("%v\n", tc.nums)
		res := maxArea(tc.nums)
		fmt.Printf("expected %d = %d\n", tc.expected, res)
	}
}
