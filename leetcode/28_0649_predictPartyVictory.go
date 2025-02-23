package main

import (
	"fmt"
	"strings"
)

// Queue
// https://leetcode.com/problems/decode-string/?envType=study-plan-v2&envId=leetcode-75

func predictPartyVictory(senate string) string {
	s := []byte(senate)
	for len(s) > 1 {
		l := s[0]
		i := 1
		for ; i < len(s); i++ {
			r := s[i]
			if l != r {
				s = append(s[:i], s[i+1:]...)
				s = s[1:]
				s = append(s, l)
				break
			}
		}
		if i == len(s) {
			break
		}

	}
	if s[0] == 'R' {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func predictPartyVictory2(senate string) string {
	var radiant, dire []int

	for idx, s := range senate {
		if string(s) == "R" {
			radiant = append(radiant, idx)
		} else {
			dire = append(dire, idx)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		// If the radiant senator comes in first, he'll block the first dire senator
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}

		// We shift the array so that we evaluate the next case
		radiant = radiant[1:]
		dire = dire[1:]
	}

	if len(radiant) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}

}

func test_predictPartyVictory() {
	testCases := []struct {
		s        string
		expected string
	}{
		{s: "RD", expected: "Radiant"},
		{s: "RDD", expected: "Dire"},
		{s: "DDRRR", expected: "Dire"},
	}
	// jkjkefjkjkefjkjkefjkjkef = jkjkefjkjkefjkjkefjkjkeff
	for _, tc := range testCases {
		fmt.Printf("%s \n", tc.s)
		res := predictPartyVictory(tc.s)
		fmt.Printf("expected %s = %s\n", tc.expected, res)
	}
	fmt.Println(strings.Repeat("*", 25))
	for _, tc := range testCases {
		fmt.Printf("%s \n", tc.s)
		res := predictPartyVictory2(tc.s)
		fmt.Printf("expected %s = %s\n", tc.expected, res)
	}
}
