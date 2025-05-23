package main

import "fmt"

/*
EX-020: Given a positive integer n, find the sum of all integers in the range [1, n] inclusive that are divisible by 3, 5, or 7. Return an integer denoting the sum of all numbers in the given range satisfying the constraint.
*/

/*
1) Understanding the problem: The problem basically gives us a number (n) and wants us to calculate the sum of all numbers from 1 to n inclusive that are divisible by either 3, 5 or 7.

2) Algorithm:
	inputs: n
	outputs: sum
	algorithm:
		- sum = 0
		- loop from 1 upwards to n
		- if i%3 == 0 || i%5 == 0 || i%7 == 0
		- then sum += i
		- return sum
3) Flowchart attached.
4) Pseudo code =>
	- sum := 0
	- for i := 1; i <= n; i++ {}
	- if i%3 == 0 || i%5 == 0 || i%7 == 0 then sum += i
	- return sum
5) Actual code below:
*/

func main() {
	n := 1000

	sum := 0

	// First way
	// for i := 1; i <= n; i++ {
	// 	if i%3 == 0 || i%5 == 0 || i%7 == 0 {
	// 		sum += i
	// 	}
	// }

	// Second way
	// for i := 1; i <= n; i++ {
	// 	if i%3 == 0 {
	// 		sum += i
	// 	}
	// }

	// for i := 1; i <= n; i++ {
	// 	if i%5 == 0 {
	// 		sum += i
	// 	}
	// }

	// for i := 1; i <= n; i++ {
	// 	if i%7 == 0 {
	// 		sum += i
	// 	}
	// }

	// for i := 1; i <= n; i++ {
	// 	if i%15 == 0 {
	// 		sum -= i
	// 	}
	// }

	// for i := 1; i <= n; i++ {
	// 	if i%21 == 0 {
	// 		sum -= i
	// 	}
	// }

	// for i := 1; i <= n; i++ {
	// 	if i%35 == 0 {
	// 		sum -= i
	// 	}
	// }

	// for i := 1; i <= n; i++ {
	// 	if i%105 == 0 {
	// 		sum += i
	// 	}
	// }

	// Third way
	multipleOf3 := 3 * (n / 3) * ((n / 3) + 1) / 2
	multipleOf5 := 5 * (n / 5) * ((n / 5) + 1) / 2
	multipleOf7 := 7 * (n / 7) * ((n / 7) + 1) / 2

	multipleOf15 := 15 * (n / 15) * ((n / 15) + 1) / 2
	multipleOf21 := 21 * (n / 21) * ((n / 21) + 1) / 2
	multipleOf35 := 35 * (n / 35) * ((n / 35) + 1) / 2

	multipleOf105 := 105 * (n / 105) * ((n / 105) + 1) / 2

	sum = multipleOf3 + multipleOf5 + multipleOf7 - multipleOf15 - multipleOf21 - multipleOf35 + multipleOf105

	fmt.Println(sum)
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 7, 10, 9, 1000
	- Edge cases:
		- 0
		- Negative numbers, but like previous questions, we assume those situations aren't going to happen
*/

/*
7) Maintenance:
	- Can I improve the performance? I guess I can. It's mostly similar to the first euler projects question, which is EX-009 in this repository; and we did the same approach there but there's actually a better approach. If we separate the conditions, like once calculate all the multiples of 3, once for 5 and once for 7 and then at last we subtract the summation with all the variations of multiples of 3*5 (15) and 3*7 (21) and 7*5 (35), we would reach the same answer I suppose. I am not sure which has the better performance here but we try it. Actually, when we're iterating through all numbers from 1 to n multiple times to find out all multiples of 3 or 5 or 7 and exclude their unions and include their multiplication; it's not going to help for performance that much! But If we use another method to find out the summation of all multiples below "n", that could help! Mathematically, we know summation of numbers 1 to n equals to n(n+1) / 2 . This has been proven. If we want to calculate summation of all numbers dividable by let's say 3, we can say that it is 3+6+9+12+15+... ; which if we factorise it, would be 3 * (1 + 2 + 3 + 4 + ...), but where's the end ? n/3 would be the end for it. So right now we know that to find out the summation of all multiplies of 3, the formula becomes 3 * ((n/3 + n/3 + 1)) / 2! Expanding it to all numbers, we can safely say that if we want the summation of all multiplies of a number "x" till n, we can calculate the endpoint = n / x and then [x * ((endpoint * (endpoint + 1))) / 2] is the final answer. This wouldn't use any loops or iteration and dramatically boosts the performance. So look at the third solution, as it does the same thing.
	- Can I make the code cleaner/more readable? With functions, yes, as we've already talked about it. Imagine writing a single function for multiplesOfXTillN, boom!
	- Document anything I change later on for making this code better.
*/

// Takeaway: Summation of multiples of "x" till "n" is x * [(n/x) * (n/x + 1)] / 2
