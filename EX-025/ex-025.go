package main

import "fmt"

/*
EX-025: Given two strings s and t, return true if t is an anagram of s, and false otherwise. (Leetcode #242)

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true

Example 2:

Input: s = "rat", t = "car"
Output: false

Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.

Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

func isAnagram(s string, t string) bool {
	freq := [26]int{}

	for _, r := range s {
		freq[r-'a']++
	}

	for _, r := range t {
		freq[r-'a']--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}

	return true
}

// Takeaway: Using maps for frequency counting is awesome. We can also use reverse functions and sorting function to solve this exercise, but the time complexity of this solution is better than them, overall. Better algorithm instead of a map, since we're dealing with alphabet is to use an array of 27 length, using it's indexes as the map keys of alphabetical characters. Because the question states that s and t consists of only lowercase english letters, this solution works better.
