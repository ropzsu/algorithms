/* https://leetcode.com/problems/longest-palindrome/#/description
Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

Example:

    Input:
    "abccccdd"
    Output:
    7

Explanation:
    One longest palindrome that can be built is "dccaccd", whose length is 7.
*/

package leetcode

func longestPalindrome(s string) int {
	maps := make(map[rune]int, len(s))
	for _, r := range s {
		maps[r]++
	}

	count := 0
	for _, v := range maps {
		if v >= 2 {
			count += v - v%2
		}
	}

	if count < len(s) {
		count++
	}
	return count
}
