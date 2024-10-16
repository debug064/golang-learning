package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(len(os.Args), os.Args)
	if len(os.Args) <= 1 {
		fmt.Println("Need test name")
		return
	}
	name := os.Args[1]
	if name == "" {
		name = "283"
	}
	if name == "mergeAlternately" {
		fmt.Println(mergeAlternately("Hello", "world"))
	} else if name == "gcdOfStrings" {
		fmt.Println(gcdOfStrings("ABABABAB", "ABAB"))
		fmt.Println(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"))
		fmt.Println(gcdOfStrings("aaaaa", "aa"))
		fmt.Println(gcdOfStrings("Hello", "world"))
		fmt.Println(gcdOfStrings("Hello", "abab"))

		fmt.Println(gcdOfStrings("aaaaa", "aaa"))
		fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	} else if name == "reverseList" {
		test_reverseList()
	} else if name == "104" {
		test_maxDepth()
	} else if name == "283" {
		test_moveZeroes()
	}
}
