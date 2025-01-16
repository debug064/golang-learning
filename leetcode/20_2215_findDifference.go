package main

import (
	"fmt"
	"strings"
)

// Hash Map / Set
// https://leetcode.com/problems/find-the-difference-of-two-arrays/description/?envType=study-plan-v2&envId=leetcode-75

func findDifference(nums1 []int, nums2 []int) [][]int {
	res := [][]int{{}, {}}
	getDif := func(nums1 []int, nums2 []int) []int {
		m := make(map[int]bool)
		for i := 0; i < len(nums1); i++ {
			m[nums1[i]] = true
		}
		diff := []int{}
		for i := 0; i < len(nums2); i++ {
			_, ok := m[nums2[i]]
			if !ok {
				diff = append(diff, nums2[i])
				m[nums2[i]] = true
			}
		}
		return diff
	}

	res[0] = getDif(nums2, nums1)
	res[1] = getDif(nums1, nums2)
	return res
}

func findDifference2(nums1 []int, nums2 []int) [][]int {
	answer := make([][]int, 2)
	answer[0] = make([]int, 0, len(nums1))
	answer[1] = make([]int, 0, len(nums2))
	nums := make(map[int]int, len(nums1)+len(nums2))
	for _, v := range nums1 {
		nums[v] = -1
	}
	for _, v := range nums2 {
		nums[v] += 2
	}
	for k, v := range nums {
		if v == -1 {
			answer[0] = append(answer[0], k)
		} else if v%2 == 0 {
			answer[1] = append(answer[1], k)
		}
	}
	return answer
}

func testFindDifference() {
	testCases := []struct {
		nums1     []int
		nums2     []int
		expected1 []int
		expected2 []int
	}{
		{nums1: []int{1, 2, 3}, nums2: []int{2, 4, 6}, expected1: []int{1, 3}, expected2: []int{4, 6}},
		{nums1: []int{1, 2, 3, 3}, nums2: []int{1, 1, 2, 2}, expected1: []int{3}, expected2: []int{}},
	}

	for _, tc := range testCases {
		fmt.Printf("[%v], [%v] \n", tc.nums1, tc.nums2)
		res := findDifference(tc.nums1, tc.nums2)
		fmt.Printf("expected [%v] [%v] = %d\n", tc.expected1, tc.expected2, res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, tc := range testCases {
		fmt.Printf("[%v], [%v] \n", tc.nums1, tc.nums2)
		res := findDifference2(tc.nums1, tc.nums2)
		fmt.Printf("expected [%v] [%v] = %d\n", tc.expected1, tc.expected2, res)
	}
}
