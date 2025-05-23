package main

import "fmt"

/*
EX-019: Given two positive integers a and b, return the number of common factors of a and b. An integer x is a common factor of a and b if x divides both a and b.
*/

/*
1) Understanding the problem: We know how to calculate GCD of two numbers. This exercise just wants all the common factors, not just the highest. We know 1 is a common factor for all numbers so we just loop from 2 upwards and count from 1.

2) Algorithm:
	inputs: a, b
	outputs: count
	algorithm:
		- loop from 2 upwards to the smaller number ( cause their common factors are surely below that)
		- if the number is divided by both a and b so count++
		- else continue
		- return count
3) Flowchart attached.
4) Pseudo code =>
	- count := 1
	- for i := 2; i <= math.Min(a, b); i++ {}
	- If a%i == 0 && b%i == 0 Then count++
	- Else continue
	- Print count
5) Actual code below:
*/

func main() {
	a := 1000
	b := 1500
	count := 1

	// First way
	// for i := 2; i <= int(math.Min(float64(a), float64(b))); i++ {
	// 	if a%i == 0 && b%i == 0 {
	// 		count++
	// 	}
	// }

	// fmt.Println(count)

	// Second way
	for b != 0 {
		a, b = b, a%b
	}

	for i := 2; i <= a; i++ {
		if a%i == 0 {
			count++
		}
	}

	fmt.Println(count)
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 12, 6 - 25, 30 - 1000, 1500
	- Edge cases:
		- Negative numbers but we assume the inputs aren't negative.
		- 0 but we again assume the input cannot be 0, but in serious situations we should consider these edge cases; as we did in previous exercises.
*/

/*
7) Maintenance:
	- Can I improve the performance? Indeed. Why do we count all the numbers from 2 to the smallest number of a and b ? Imagine one of the number is 5 million, that's not optimal. Instead, we know that we can calculate their GCD, and it's obvious common factors are always beneath that number! So instead of the smallest number, we calculate the GCD and loop upwards till we reach there.
	- Can I make the code cleaner/more readable? Yeah, with functions we could separate the GCD function totally and that's the thing we're going to do in leetcode when we're going to submit our solution, but for now since we haven't studied functions, we're going to do like above. Also, mathematically, if any factor divides a and b exactly, it would divide their GCD exactly too; so we can only check gcd%i and it's going to work and is mathematically correct. Furthermore, if we only use gcd, we don't even need a,b anymore so we don't need those excessive x and y variables.
	- Document anything I change later on for making this code better.
*/

// Takeaway: All common factors of two numbers are below their GCD, since their greatest common divisor is that number! Also, any number that divides two numbers a and b exactly, also divides their GCD exactly too.
