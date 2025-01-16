package main

import "fmt"

// Stack
// https://leetcode.com/problems/removing-stars-from-a-string/description/?envType=study-plan-v2&envId=leetcode-75

func removeStars(s string) string {
	var chars []rune
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			chars = chars[:max(0, len(chars)-1)]
		} else {
			chars = append(chars, rune(s[i]))
		}
	}
	return string(chars)
}

func testRemoveStars() {
	testCases := []struct {
		s        string
		expected string
	}{
		{s: "leet**cod*e", expected: "lecoe"},
		{s: "erase*****", expected: ""},
	}

	for _, tc := range testCases {
		fmt.Printf("%s \n", tc.s)
		res := removeStars(tc.s)
		fmt.Printf("expected %s = %s\n", tc.expected, res)
	}
}
