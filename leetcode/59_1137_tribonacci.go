package main

import (
	"fmt"
	"strings"
)

// DP - 1D
// https://leetcode.com/problems/n-th-tribonacci-number/description/?envType=study-plan-v2&envId=leetcode-75

func tribonacci(n int) int {
	T := []int{0, 1, 1}

	if n < len(T) {
		return T[n]
	}

	res := 0
	for i := 3; i <= n; i++ {
		res = T[0] + T[1] + T[2]
		T[0] = T[1]
		T[1] = T[2]
		T[2] = res
	}
	return res
}

func testTribonacci() {
	testCases := []struct {
		n        int
		expected int
	}{
		{n: 4, expected: 4},
		{n: 25, expected: 1389537},
	}

	for _, tc := range testCases {
		fmt.Printf("[%d] \n", tc.n)
		res := tribonacci(tc.n)
		fmt.Printf("expected [%d] = %d\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
