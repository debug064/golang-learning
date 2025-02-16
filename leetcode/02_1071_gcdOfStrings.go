package main

import (
	"fmt"
	"strings"
)

func gcdOfStringsV(str1 string, str2 string) string {

	base := str2
	size := len(base)
	for i := 1; i < size/2+1; i++ {
		if size%i != 0 {
			continue
		}
		s := base[:i]
		if strings.Count(base, s) == size/i {
			base = s
			break
		}
	}

	if strings.Count(str1, base) == len(str1)/len(base) {
		return base
	}

	return ""
}

func gcdOfStrings(str1 string, str2 string) string {

	size1 := len(str1)
	size2 := len(str2)

	gcd := 1
	for i := 1; i <= size1 && i <= size2; i++ {
		if size1%i == 0 && size2%i == 0 {
			gcd = i
		}
	}

	if strings.Count(str2, str2[:gcd]) != size2/gcd {
		return ""
	}

	if strings.Count(str1, str2[:gcd]) == size1/gcd {
		return str2[:gcd]
	}

	return ""
}

func test_gcdOfStrings() {
	fmt.Println(gcdOfStrings("ABABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"))
	fmt.Println(gcdOfStrings("aaaaa", "aa"))
	fmt.Println(gcdOfStrings("Hello", "world"))
	fmt.Println(gcdOfStrings("Hello", "abab"))

	fmt.Println(gcdOfStrings("aaaaa", "aaa"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))

	fmt.Println(gcdOfStringsV("aaaaa", "aaa"))
	fmt.Println(gcdOfStringsV("ABABAB", "ABAB"))
}
