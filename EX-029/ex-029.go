package main

import "fmt"

/*
EX-029: Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums. (Leetcode #448)

Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]

Example 2:

Input: nums = [1,1]
Output: [2]

Constraints:

n == nums.length
1 <= n <= 105
1 <= nums[i] <= n

Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
*/

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := freq[i]; ok == false {
			result = append(result, i)
		}
	}

	return result
}

// Takeaway: Actually, freq[i] return the zero value of the map's value type by default if the element does not exist, and if it exists, returns the actual value of that key.
