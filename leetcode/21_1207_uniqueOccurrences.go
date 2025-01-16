package main

import "fmt"

// Hash Map / Set
// https://leetcode.com/problems/unique-number-of-occurrences/description/?envType=study-plan-v2&envId=leetcode-75

func uniqueOccurrences(arr []int) bool {
	nums := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		nums[arr[i]] += 1
	}
	cnts := make(map[int]int)
	for _, v := range nums {
		_, ok := cnts[v]
		if ok {
			return false
		}
		cnts[v] += 1
	}
	return true
}

func testUniqueOccurrences() {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{nums: []int{1, 2, 2, 1, 1, 3}, expected: true},
		{nums: []int{1, 2}, expected: false},
		{nums: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, expected: true},
	}

	for _, tc := range testCases {
		fmt.Printf("%v\n", tc.nums)
		res := uniqueOccurrences(tc.nums)
		fmt.Printf("expected %t = %t\n", tc.expected, res)
	}
}
