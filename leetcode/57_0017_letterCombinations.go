package main

import (
	"fmt"
	"strconv"
)

// Backtracking
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/?envType=study-plan-v2&envId=leetcode-75

func letterCombinations(digits string) []string {
	phone := map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	res := []string{}

	var comb func(nums string)
	comb = func(nums string) {
		if len(nums) == 0 {
			return
		}
		if len(nums) == 1 {
			n, _ := strconv.Atoi(nums)
			for _, c := range phone[n] {
				res = append(res, string(c))
			}
			return
		}
		comb(nums[1:])
		tmp := []string{}
		n, _ := strconv.Atoi(string(nums[0]))
		for _, v := range phone[n] {
			for _, s := range res {
				tmp = append(tmp, string(v)+s)
			}
		}
		res = tmp
	}
	comb(digits)
	return res
}

func testLetterCombinations() {
	testCases := []struct {
		digits   string
		expected []string
	}{
		{digits: "23", expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{digits: "", expected: []string{""}},
		{digits: "2", expected: []string{"a", "b", "c"}},
	}

	for _, tc := range testCases {
		fmt.Printf("%s \n", tc.digits)
		res := letterCombinations(tc.digits)
		fmt.Printf("expected %v = %v\n", tc.expected, res)
	}
}
