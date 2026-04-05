// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

package main

import (
	"fmt"
	"strings"
)

func theLongestWord(str string) string {
	var longestWord string
	words := strings.Fields(str)

	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	return longestWord
}

func main() {
	fmt.Println(theLongestWord("I like my new lifestyle"))
}
