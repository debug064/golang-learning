package main

import (
	"fmt"
	"slices"
	"strings"
)

// Monotonic Stack
// https://leetcode.com/problems/daily-temperatures/?envType=study-plan-v2&envId=leetcode-75

func dailyTemperatures(temperatures []int) []int {

	type Pair struct {
		First  int
		Second int
	}

	res := []int{}
	mstack := []Pair{}

	for i, v := range slices.Backward(temperatures) {
		//fmt.Println(i, ":", v)
		for len(mstack) > 0 && v >= mstack[len(mstack)-1].First {
			mstack = mstack[:len(mstack)-1]
		}
		if len(mstack) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, mstack[len(mstack)-1].Second-i)
		}
		mstack = append(mstack, Pair{First: v, Second: i})

	}
	slices.Reverse(res)
	return res
}

func testDailyTemperatures() {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}, expected: []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0}},
		{nums: []int{73, 74, 75, 71, 69, 72, 76, 73}, expected: []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{nums: []int{30, 40, 50, 60}, expected: []int{1, 1, 1, 0}},
		{nums: []int{30, 60, 90}, expected: []int{1, 1, 0}},
	}

	for _, tc := range testCases {
		fmt.Printf("%v \n", tc.nums)
		res := dailyTemperatures(tc.nums)
		fmt.Printf("expected %v = %v\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))

}
