package main

import "fmt"

// Stack
// https://leetcode.com/problems/asteroid-collision/description/?envType=study-plan-v2&envId=leetcode-75

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	for _, v := range asteroids {
		for len(res) > 0 && res[len(res)-1] > 0 && v < 0 {
			if res[len(res)-1] < -v {
				res = res[:len(res)-1]
			} else if res[len(res)-1] == -v {
				res = res[:len(res)-1]
				v = 0
			} else {
				v = 0
				break
			}
		}
		if v != 0 {
			res = append(res, v)
		}
	}
	return res
}

func testAsteroidCollision() {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{5, 10, -5}, expected: []int{5, 10}},
		{nums: []int{8, -8}, expected: []int{}},
		{nums: []int{10, 2, -5}, expected: []int{10}},
	}

	for _, tc := range testCases {
		fmt.Printf("%v \n", tc.nums)
		res := asteroidCollision(tc.nums)
		fmt.Printf("expected %v = %v\n", tc.expected, res)
	}
}
