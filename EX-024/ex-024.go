package main

import "fmt"

/*
EX-024: Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct. (Leetcode #217)

Example 1:

Input: nums = [1,2,3,1]
Output: true
Explanation: The element 1 occurs at the indices 0 and 3.

Example 2:

Input: nums = [1,2,3,4]
Output: false
Explanation: All elements are distinct.

Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	for _, v := range freq {
		if v >= 2 {
			return true
		}
	}
	return false
}

// Takeaway: We can solve this problem using sorting algorithms too.
