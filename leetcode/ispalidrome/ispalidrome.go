// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// LeetCode 125: Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/
//
// Problem Statement:
// A phrase is a palindrome if, after converting all uppercase letters into lowercase
// letters and removing all non-alphanumeric characters, it reads the same forward and
// backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Constraints:
// - 1 <= s.length <= 2 * 10^5
// - s consists only of printable ASCII characters.

package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Go approach using standard library
// Time Complexity: O(n), Space Complexity: O(n)
func isPalindrome(s string) bool {
	var builder strings.Builder
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			builder.WriteRune(unicode.ToLower(c))
		}
	}
	str := builder.String()

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

// // Two-pointer approach: O(n) time, O(1) space
// func isPalindrome(s string) bool {
// 	left := 0
// 	right := len(s) - 1

// 	for left < right {
// 		if !isAlphanumeric(s[left]) {
// 			left++
// 			continue
// 		}
// 		if !isAlphanumeric(s[right]) {
// 			right--
// 			continue
// 		}
// 		if toLower(s[left]) != toLower(s[right]) {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }

// // Checks if character is alphanumeric (letter or digit)
// func isAlphanumeric(c byte) bool {
// 	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
// }

// // Converts uppercase ASCII letter to lowercase using ASCII offset (32)
// func toLower(c byte) byte {
// 	if c >= 'A' && c <= 'Z' {
// 		return c + 32
// 	}
// 	return c
// }
