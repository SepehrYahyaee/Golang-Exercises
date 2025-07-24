package main

import "fmt"

/*
EX-022: Given a non-empty array of integers nums, every element appears twice except for one. Find that single one. You must implement a solution with a linear runtime complexity and use only constant extra space. (Leetcode #136)

Example 1:

Input: nums = [2,2,1]

Output: 1

Example 2:

Input: nums = [4,1,2,1,2]

Output: 4

Example 3:

Input: nums = [1]

Output: 1

Constraints:
1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears only once.
*/

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}

func singleNumber(nums []int) int {
	var result int
	for _, v := range nums {
		result ^= v
	}
	return result
}

// Takeaway: Using XOR's is pretty fine with calculating these frequency stuff, since XOR of 0 with anything is that number itself, and XOR of two numbers equal to each other is 0. We can also brute-force our way and do nested loops but that's not optimal, obviously.
