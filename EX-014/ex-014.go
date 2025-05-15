package main

import "fmt"

/*
EX-014: A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 times 99. Find the largest palindrome made from the product of two 3-digit numbers.
*/

/*
1) Understanding the problem: We solved a palinrome exercise in the last question, so we know what it means. Now, we got to find the largest palindrome number that is the product of two 3-digit integers. Multiplying two 3-digit numbers is going to result in a number having 5 or 6 digits since 100*100=10000 and 999*999 is going to be 998,001 which has 6 digits. So the palindrome we're looking for isn't out of this range. Since we're looking for the largest one, it's better to start counting from the upper end, doing 999*999 and 999*998 and so on to reach the target, and then 998 times 999 and so on.

2) Algorithm:
	inputs: -
	outputs: largest palindrome that is made up by the product of two 3-digit numbers
	algorithm:
		- loop from 999 down
		- having a nested loop from 999 down inside the outer loop
		- calculate the product of them
		- calculate if the result is palindrome or not
3) Flowchart attached.
4) Pseudo code =>
	- max = 0
	- for i = 999; i >= 100; i-- { for j = 999; j >= 100; j-- {} }
	- Check isPalindrome inside the inner loop (copy of the last question)
	- if palindrome, update the max and move on to the next candidate.
	- At last when all loops ended, print out the max
5) Actual code below:
*/

func main() {
	max := 0

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j

			if product < max {
				break
			}

			if product%10 == 0 {
				continue
			}

			temp, reversed := product, 0

			for temp > reversed {
				reversed = 10*reversed + temp%10
				temp /= 10
			}

			if (temp == reversed || temp == reversed/10) && product > max {
				max = product
			}
		}
	}

	fmt.Println(max)
}

/*
6) Testing: it was a calculation process, no testing needed!
*/

/*
7) Maintenance:
	- Can I improve the performance? Yes, a few optimizations would work. We can check for product not being smaller than max in the starting of the inner loop; that shows if the product is going to be smaller than max, then what's the point in checking it. You can safely ignore the rest of that loop since they are only decreasing and all the others are going to be smaller numbers! (Simple mathematics) [Line 37 to 39]
	- Can I make the code cleaner/more readable? Yes, isPalindrome is not needed and we can safely check the condition inside the if/statement later on, but I'm not sure if it's more readable or not, but it's definitely cleaner. Also inlining the two reversed and temp variable in one declaration using multiple declaration style (a, b := 1, 2 for example) is better for readability.
	- Document anything I change later on for making this code better.
*/

// Takeaway: Nothing special; just that when we are looping downwards from a number, obviously the smaller numbers and their products is going to be smaller than the previous ones and so if our condition is to check one large thing, like here we got to find the largest palindrome; we could easily break the loop when it's past the point we want. This simple mathematics stuff, which I've missed at first, is going to help the performance much more than it'll look. We'll learn how to use strings functions to solve these kind of palindrome questions too but that solutions wouldn't be better than these ones anyway.
