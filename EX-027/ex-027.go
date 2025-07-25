package main

import "fmt"

/*
EX-027: Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1. (Leetcode #387)

Example 1:

Input: s = "leetcode"
Output: 0
Explanation: The character 'l' at index 0 is the first character that does not occur at any other index.

Example 2:

Input: s = "loveleetcode"
Output: 2

Example 3:

Input: s = "aabb"
Output: -1

Constraints:

1 <= s.length <= 105
s consists of only lowercase English letters.
*/

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}

func firstUniqChar(s string) int {
	freq := [26]uint{}

	for _, v := range s {
		freq[v-'a']++
	}

	for i, r := range s {
		if freq[r-'a'] == 1 {
			return i
		}
	}

	return -1
}

// Takeaway: Nothing special.
