package main

import (
	"fmt"
	"math"
)

/*
EX-018: Given a positive integer n, return the smallest positive integer that is a multiple of both 2 and n.
*/

/*
1) Understanding the problem: It wants the LCM of two numbers but those numbers are always 2 and n. We know that the LCM of two numbers is their multiplication divided by their GCD, so the answer to this question is always 2*n / GCD(2, n). Now GCD(2, n) depends on the number n being even or odd. If the n is even so their GCD(2, n) is 2 and the overall answer is "n" itself. If the number n is odd, then they have no common divisor and their GCD is 1 which makes the whole equation 2*n. Putting it simply, the answer is either 2*n (when n is odd) or n (when n is even).

2) Algorithm:
	inputs: n
	outputs: lcm(2, n)
	algorithm:
		- If n%2 == 0
		- Then answer is n
		- Else
		- Answer is 2*n
3) Flowchart attached.
4) Pseudo code =>
	- If n&1 == 0 { return n }
	- return 2*n
5) Actual code below:
*/

func main() {
	n := int(math.Abs(float64(1)))

	if n < 0 {
		fmt.Println("n should not be zero.")
	} else {
		if n&1 == 0 {
			fmt.Println(n)
		} else {
			fmt.Println(2 * n)
		}
	}
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 150, 1
	- Edge cases:
		- Negative numbers, LCM won't care about those.
*/

/*
7) Maintenance:
	- Can I improve the performance? Using n&1 instead of n%2, though the compiler might do it itself, might make the performance better.
	- Can I make the code cleaner/more readable? I don't know what is more readable than this. This version looks boring but learning and using functions would make it clearly readable.
	- Document anything I change later on for making this code better.
*/

// Takeaway: Nothing.
