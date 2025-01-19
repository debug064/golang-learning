package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Stack
// https://leetcode.com/problems/decode-string/?envType=study-plan-v2&envId=leetcode-75

func decodeString(s string) string {
	last := func(b int) int {
		n := 0
		for i := b; i < len(s); i++ {
			if s[i] == '[' {
				n += 1
			} else if s[i] == ']' {
				n -= 1
				if n == 0 {
					return i - b
				}
			}
		}
		return len(s) - b
	}
	num := 1
	begin := 0
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			break
		}
		if s[i] == '[' {
			end := last(i)
			str := decodeString(s[i+1 : i+end])
			i += end
			begin = i + 1
			sb.WriteString(strings.Repeat(str, num))
			num = 1
			continue
		}
		n, e := strconv.Atoi(s[begin : i+1])
		if e == nil {
			num = n
		} else {
			sb.WriteByte(s[i])
			begin = i + 1
		}
		//fmt.Printf("%d %s\n", num, sb.String())
	}
	return sb.String()
}

func testDecodeString() {
	testCases := []struct {
		s        string
		expected string
	}{
		{s: "3ss[a]2[bc]", expected: "aaabcbc"},
		{s: "4[2[jk]e1[f]]", expected: "jkjkefjkjkefjkjkefjkjkef"},
		//{s: "3[z]2[2[y]pq4[2[jk]e1[f]]]ef", expected: "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"},
		{s: "3[a2[c]]", expected: "accaccacc"},
		{s: "2[abc]3[cd]ef", expected: "abcabccdcdcdef"},
	}
	// jkjkefjkjkefjkjkefjkjkef = jkjkefjkjkefjkjkefjkjkeff
	for _, tc := range testCases {
		fmt.Printf("%s \n", tc.s)
		res := decodeString(tc.s)
		fmt.Printf("expected %s = %s\n", tc.expected, res)
	}
}
