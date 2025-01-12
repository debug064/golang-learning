package main

import "fmt"

// Sliding Window
// https://leetcode.com/problems/max-consecutive-ones-iii/description/?envType=study-plan-v2&envId=leetcode-75

func longestOnes(nums []int, k int) int {
	m := 0
	j := 0
	zero := 0
	//t := 0
	for i, n := range nums {
		//fmt.Printf("%d\n", n)
		if n == 1 {
			//t++
			if i-j+1 > m {
				m = i - j + 1
			}
			continue
		}
		zero++
		if zero > k {
			for ; j < len(nums); j++ {
				if nums[j] == 0 {
					zero--
					j++
					break
				}
			}
		}
		if i-j+1 > m {
			m = i - j + 1
		}
	}
	return m
}

func testLongestOnes() {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, k: 2, expected: 6},
		{nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 3, expected: 10},
	}

	for _, tc := range testCases {
		fmt.Printf("%v - %d\n", tc.nums, tc.k)
		res := longestOnes(tc.nums, tc.k)
		fmt.Printf("expected %d = %d\n", tc.expected, res)
	}
}
