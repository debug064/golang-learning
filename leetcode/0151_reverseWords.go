package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/reverse-words-in-a-string/description/?envType=study-plan-v2&envId=leetcode-75

func reverseWords(s string) string {
	runes := []rune(s)
	space := 0
	var sb strings.Builder

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			if space == 1 {
				space = 2
			}
			continue
		} else if space == 2 {
			sb.WriteRune(' ')
		}
		sb.WriteRune(runes[i])
		space = 1
	}
	runes = []rune(sb.String())
	l, r := 0, 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			for ; l < r; l, r = l+1, r-1 {
				runes[l], runes[r] = runes[r], runes[l]
			}
			l = i + 1
			r = l
			continue
		}
		r = i
	}
	for ; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}

	return string(runes)
}

func reverseWords2(s string) string {
	runes := []rune(s)
	l := 0
	spaces := true
	for r := l; r < len(runes); r++ {
		if runes[r] == ' ' {
			if spaces {
				continue
			}
			spaces = true
		} else {
			spaces = false
		}
		runes[l] = runes[r]

		l++
	}
	if runes[l-1] == ' ' {
		l--
	}
	size := l
	l = 0

	for r := size - 1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}

	l, r := 0, 0
	for i := 0; i < size; i++ {
		if runes[i] == ' ' {
			for ; l < r; l, r = l+1, r-1 {
				runes[l], runes[r] = runes[r], runes[l]
			}
			l = i + 1
			r = l
			continue
		}
		r = i
	}
	for ; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}

	return string(runes[:size])
}

func reverseWords3(s string) string {
	var sb strings.Builder
	res := strings.Split(s, " ")
	added := false
	for i := len(res) - 1; i >= 0; i-- {
		if len(res[i]) == 0 {
			continue
		}
		if added {
			sb.WriteRune(' ')
			added = false
		}
		sb.WriteString(res[i])
		added = true

	}

	return sb.String()
}

func reverseWords4(s string) string {
	var sb strings.Builder
	i := len(s) - 1
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		end := i
		if i < 0 {
			break
		}
		if sb.Len() > 0 {
			sb.WriteString(" ")
		}
		for i >= 0 && s[i] != ' ' {
			i--
		}
		sb.WriteString(s[i+1 : end+1])
	}
	return sb.String()
}

type TestData151 struct {
	str string
	res string
}

func test_reverseWords() {
	var tests = []TestData151{
		{str: "the sky is blue", res: "blue is sky the"},
		{str: "  hello world  ", res: "world hello"},
		{str: "a good   example", res: "example good a"},
	}

	var res string
	for _, val := range tests {
		fmt.Println(val)
		res = reverseWords(val.str)
		fmt.Printf("expected '%s' = '%s'\n", val.res, res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, val := range tests {
		fmt.Println(val)
		res = reverseWords2(val.str)
		fmt.Printf("expected '%s' = '%s'\n", val.res, res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, val := range tests {
		fmt.Println(val)
		res = reverseWords3(val.str)
		fmt.Printf("expected '%s' = '%s'\n", val.res, res)
	}
}
