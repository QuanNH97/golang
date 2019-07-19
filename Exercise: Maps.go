package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, count := range a {
		m[count]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
