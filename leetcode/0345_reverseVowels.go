package main

import (
	"fmt"
	"strings"
	"unicode"
)

// https://leetcode.com/problems/reverse-vowels-of-a-string/?envType=study-plan-v2&envId=leetcode-75

func reverseVowels(s string) string {
	l := 0
	r := len(s) - 1
	runes := []rune(s)
	vowels := "aeiou"
	for ; l < r; l, r = l+1, r-1 {
		for ; l < r; l += 1 {
			c := unicode.ToLower(runes[l])
			cl := strings.IndexRune(vowels, c)
			if cl != -1 {
				break
			}

		}
		for ; l < r; r -= 1 {
			c := unicode.ToLower(runes[r])
			cr := strings.IndexRune(vowels, c)
			if cr != -1 {
				break
			}
		}

		if l < r {
			runes[l], runes[r] = runes[r], runes[l]
		}
	}
	return string(runes)
}

func reverseVowels2(s string) string {
	l := 0
	r := len(s) - 1
	runes := []rune(s)
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for ; l < r; l, r = l+1, r-1 {
		for ; l < r; l += 1 {
			c := unicode.ToLower(runes[l])
			_, ok := vowels[c]
			if ok {
				break
			}

		}
		for ; l < r; r -= 1 {
			c := unicode.ToLower(runes[r])
			_, ok := vowels[c]
			if ok {
				break
			}
		}

		if l < r {
			runes[l], runes[r] = runes[r], runes[l]
		}
	}
	return string(runes)
}

type TestData345 struct {
	str string
	res string
}

func test_reverseVowels() {
	var tests = []TestData345{
		{str: "IceCreAm", res: "AceCreIm"},
		{str: "leetcode", res: "leotcede"},
	}

	var res string
	for _, val := range tests {
		fmt.Println(val)
		res = reverseVowels(val.str)
		fmt.Printf("expected %s = %s\n", val.res, res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, val := range tests {
		fmt.Println(val)
		res = reverseVowels2(val.str)
		fmt.Printf("expected %s = %s\n", val.res, res)
	}
}
