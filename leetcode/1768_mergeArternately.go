package main

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	sb.Grow(len(word1) + len(word2))
	i := 0
	//j := 0
	for ; i < len(word1) && i < len(word2); i++ {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[i])
	}
	if len(word1) > len(word2) {
		sb.WriteString(word1[i:])
	} else if len(word1) < len(word2) {
		sb.WriteString(word2[i:])
	}
	return sb.String()
}
