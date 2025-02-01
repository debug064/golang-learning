package main

import (
	"fmt"
	"strings"
)

// DP - 1D
// https://leetcode.com/problems/min-cost-climbing-stairs/description/?envType=study-plan-v2&envId=leetcode-75

func minCostClimbingStairs(cost []int) int {
	pr := make([]int, len(cost))
	pr[0] = cost[0]
	pr[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		pr[i] = cost[i] + min(pr[i-1], pr[i-2])
	}
	res := min(pr[len(pr)-1], pr[len(pr)-2])
	return res
}

func testMinCostClimbingStairs() {
	testCases := []struct {
		n        []int
		expected int
	}{
		{n: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, expected: 6},
		{n: []int{0, 1, 1, 1}, expected: 1},
		{n: []int{10, 15}, expected: 10},
		{n: []int{10, 15, 20}, expected: 15},
	}

	for _, tc := range testCases {
		fmt.Printf("[%v] \n", tc.n)
		res := minCostClimbingStairs(tc.n)
		fmt.Printf("expected [%d] = %d\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
}
