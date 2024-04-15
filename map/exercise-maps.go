//go:build ignore

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	for _, word := range words {
		result[word]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
