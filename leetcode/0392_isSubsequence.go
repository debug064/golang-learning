package main

import "fmt"

// https://leetcode.com/problems/is-subsequence/?envType=study-plan-v2&envId=leetcode-75

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	iS := 0
	iT := 0
	for iT < len(t) {
		if s[iS] == t[iT] {
			iS++
			if iS == len(s) {
				return true
			}
		}
		iT++

	}
	return false
}

func testIsSubsequence() {
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "abc", t: "ahbgdc", expected: true},
		{s: "axc", t: "ahbgdc", expected: false},
		{s: "", t: "ahbgdc", expected: true},
	}

	for _, tc := range testCases {
		fmt.Printf("%s - %s\n", tc.s, tc.t)
		res := isSubsequence(tc.s, tc.t)
		fmt.Printf("expected %t = %t\n", tc.expected, res)
	}
}
