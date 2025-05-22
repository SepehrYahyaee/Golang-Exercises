package main

import (
	"fmt"
	"math"
)

/*
EX-016: Write a program to calculate the least common multiple of two numbers (LCM).
*/

/*
1) Understanding the problem: Last problem was about GCD, greatest common divisor, one of the famous properties we often calculate and is used in many other algorithms and places. This one, is about LCM, least common multiple of two numbers that we're going to calculate. The first thing I can think of, is that LCM of two numbers is at most the multiplication of them (unless they are both the same). So let's say we got a and b as inputs, their LCM is below the a*b multiplication. So if we loop down from that value, and check if each number is dividable by both a and b, so we find the least common multiple of both of them.

2) Algorithm:
	inputs: a, b
	outputs: lcm
	algorithm:
		- calculate a*b
		- loop from a*b downwards to 0
		- check if a number is divideable by both a and b
		- if it was, count it as a potential LCM till you find the least of them
3) Flowchart attached.
4) Pseudo code =>
	- a, b = ...
	- lcm = 0
	- for i := a*b; i > 0; i-- { if i%a == 0 && i%b == 0 then lcm = i }
	- print out lcm
5) Actual code below:
*/

func main() {
	a := int(math.Abs(float64(4)))
	b := int(math.Abs(float64(6)))

	// First way
	//
	// lcm := 0

	// if a == 0 || b == 0 {
	// 	fmt.Println("a or b cannot be 0")
	// } else if a == 1 || b == 1 {
	// 	fmt.Println(a * b)
	// } else {
	// 	for i := a * b; i > 0; i-- {
	// 		if i%a == 0 && i%b == 0 {
	// 			lcm = i
	// 		}
	// 	}

	// 	fmt.Println(lcm)
	// }

	// Second way
	multiplication := a * b

	for b != 0 {
		a, b = b, a%b
	}

	fmt.Println(multiplication / a)
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 4, 6 - 12, 24 - 1, 24 - 12048292, 1243898
	- Edge cases:
		- 0 which is handled by an if statement
		- Negative numbers which is handles by math.Abs function
*/

/*
7) Maintenance:
	- Can I improve the performance? For the first solution, I can check if any of the numbers are 1 and if they are, so the answer is a*b already. But I can't seem to find any other performance boosts for that matter. But this whole approach isn't efficient for big numbers, since it's going to loop downwards from their multiplication! The second solution using the GCD of two numbers to calculate the LCM is the best approach in terms of performance. The logic is simple: You already know that at most, the LCM is a*b, but what if the two numbers have common factors within each other that both of them divides to ? So we gotta divide the a*b to their greatest common divisor (GCD) that would remove all the common factors. Any smaller common factor is already available in the GCD, so if we divide their multiplication by the GCD once, the shared parts are removed and only the common parts remains.
	- Can I make the code cleaner/more readable? Again, the functions can make the code more readable but for now, it's in it's most readable state.
	- Document anything I change later on for making this code better.
*/

// Takeaway: LCM is tightly coupled with GCD and the two is used in many algorithm and it's better to know their efficient algorithms. As always, mathematics and number theory is fascinating, at least to me!
