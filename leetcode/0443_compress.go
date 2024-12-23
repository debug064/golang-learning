package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/string-compression/?envType=study-plan-v2&envId=leetcode-75
func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	count := 1
	curr := chars[0]
	j := 0

	finish := func() {
		chars[j] = curr
		j += 1
		if count > 1 {
			for _, val := range fmt.Sprintf("%d", count) {
				chars[j] = byte(val)
				j += 1
			}
		}
	}

	for i := 1; i < len(chars); i++ {
		if curr == chars[i] {
			count++
			continue
		}

		finish()
		curr = chars[i]
		count = 1
	}
	finish()
	return j
}

func testCompress() {
	testCases := []struct {
		charsIn  []byte
		expected int
		charsOut []byte
	}{
		{charsIn: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, expected: 6, charsOut: []byte{'a', '2', 'b', '2', 'c', '3'}},
		{charsIn: []byte{'a'}, expected: 1, charsOut: []byte{'a'}},
		{charsIn: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, expected: 4, charsOut: []byte{'a', 'b', '1', '2'}},
	}

	for _, val := range testCases {
		fmt.Println(string(val.charsIn))
		res := compress(val.charsIn)
		fmt.Printf("expected %d and %v  =  %d and %v\n", val.expected, string(val.charsOut), res, string(val.charsIn))
		fmt.Println(strings.Repeat("*", 25))
	}
}
