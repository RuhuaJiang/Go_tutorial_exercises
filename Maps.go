package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	stringSlice := strings.Split(s, " ")
	m := make(map[string]int)
	for _, v := range stringSlice {
		m[v]++
	}
	return m
}
func main() {
	wc.Test(WordCount)
}
