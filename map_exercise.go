package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	words_map := make(map[string]int)
	for _, value := range(words) {
		words_map[value] = words_map[value] + 1
	}
	return words_map
}

func main() {
	wc.Test(WordCount)
}
