// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// LeetCode 125: Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/
//
// A phrase is a palindrome if, after converting all uppercase letters into lowercase
// letters and removing all non-alphanumeric characters, it reads the same forward and
// backward. Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//
// Example:
//   Input: s = "A man, a plan, a canal: Panama"
//   Output: true
//   Explanation: "amanaplanacanalpanama" is a palindrome.
//
// Constraints:
//   - 1 <= s.length <= 2 * 10^5
//   - s consists only of printable ASCII characters.
//
// Time: O(n) - each pointer moves at most n steps total
// Space: O(1) - only two indices, no extra buffer

package main

import "fmt"

// Two-pointer approach: shrink from both ends, skip non-alphanumerics
func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

// Checks if character is alphanumeric (letter or digit)
func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

// Converts uppercase ASCII letter to lowercase using ASCII offset (32)
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
}

/*
Opposite-ends Two Pointers: s = "A man, a plan, a canal: Panama"

  L                                         R
  A | ' ' | m | a | n | ... | n | a | m | a
  → both alphanumeric, toLower('A')==toLower('a') → L++, R--

  Skip non-alphanumerics (spaces, commas, colon) on either side,
  compare lowercased bytes. Mismatch → not a palindrome.

Why opposite-ends:
  - Palindrome check is symmetric: s[i] must equal s[n-1-i]
  - Converging pointers do one pass, O(1) extra space
  - Skipping non-alphanumerics is built into the pointer advance,
    no pre-filtering / no string copy needed

Alternative (standard library, O(n) space): build a filtered+lowercased
copy via strings.Builder + unicode.IsLetter/IsDigit, then compare halves.
Clean and concise, but allocates a new string — not in the spirit of the
two-pointer pattern:

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
*/
