package main

// Sliding Window
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/?envType=study-plan-v2&envId=leetcode-75

import (
	"fmt"
	"strings"
)

func maxVowels(s string, k int) int {
	vowels := "aeiou"
	m1 := 0
	m2 := 0
	for i, j := 0, 0; i < len(s); i++ {
		if strings.IndexByte(vowels, s[i]) != -1 {
			m2++
		}
		if i+1-j > k {
			if strings.IndexByte(vowels, s[j]) != -1 {
				m2--
			}
			j++
		}
		if m2 > m1 {
			m1 = m2
		}
	}
	return m1
}

func maxVowels2(s string, k int) int {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	m1 := 0
	m2 := 0
	for i, j := 0, 0; i < len(s); i++ {
		_, ok := vowels[rune(s[i])]
		if ok {
			m2++
		}
		if i+1-j > k {
			_, ok = vowels[rune(s[j])]
			if ok {
				m2--
			}
			j++
		}
		if m2 > m1 {
			m1 = m2
		}
	}
	return m1
}

func maxVowels3(s string, k int) int {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}
	m1 := 0
	m2 := 0
	for i, j := 0, 0; i < len(s); i++ {
		if isVowel(s[i]) {
			m2++
		}
		if i+1-j > k {
			if isVowel(s[j]) {
				m2--
			}
			j++
		}
		if m2 > m1 {
			m1 = m2
		}
	}
	return m1
}

func testMaxVowels() {
	testCases := []struct {
		s        string
		k        int
		expected int
	}{
		{s: "abciiidef", k: 3, expected: 3},
		{s: "aeiou", k: 2, expected: 2},
		{s: "leetcode", k: 3, expected: 2},
	}

	for _, tc := range testCases {
		fmt.Printf("%s - %d\n", tc.s, tc.k)
		res := maxVowels(tc.s, tc.k)
		fmt.Printf("expected %d = %d\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, tc := range testCases {
		fmt.Printf("%s - %d\n", tc.s, tc.k)
		res := maxVowels2(tc.s, tc.k)
		fmt.Printf("expected %d = %d\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, tc := range testCases {
		fmt.Printf("%s - %d\n", tc.s, tc.k)
		res := maxVowels3(tc.s, tc.k)
		fmt.Printf("expected %d = %d\n", tc.expected, res)
	}
}
