package main

import "fmt"

// https://leetcode.com/problems/max-number-of-k-sum-pairs/?envType=study-plan-v2&envId=leetcode-75

func maxOperations(nums []int, k int) int {
	m := make(map[int]int)
	count := 0
	// for _, v := range nums {
	// 	//fmt.Printf("%d\n", v)
	// 	d := k - v
	// 	_, ok := m[d]
	// 	if ok {
	// 		m[d] -= 1
	// 		count++
	// 		if m[d] == 0 {
	// 			delete(m, d)
	// 		}
	// 	} else {
	// 		m[v] += 1
	// 	}

	// }
	for _, v := range nums {
		d := k - v
		if m[d] > 0 {
			m[d]--
			count++
		} else {
			m[v]++
		}
	}
	return count
}

func testMaxOperations() {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{nums: []int{1, 2, 3, 4}, k: 5, expected: 2},
		{nums: []int{3, 1, 3, 4, 3}, k: 6, expected: 1},
	}

	for _, tc := range testCases {
		fmt.Printf("%v - %d\n", tc.nums, tc.k)
		res := maxOperations(tc.nums, tc.k)
		fmt.Printf("expected %d = %d\n", tc.expected, res)
	}
}
