package main

import (
	"fmt"
	"math"
)

/*
EX-007: Program to calculate the absolute amount of a given number as an input.
*/

/*
1) Understanding the problem: We are getting a number as an input, and we have to return the absolute value of it.
2) Algorithm:
	inputs: n (float)
	outputs: a (float)
	algorithm:
		- Take number
		- Calculate it's absolute, change the sign to positive (a)
		- Print out a
3) Flowchart attached.
4) Pseudo code =>
	- n
	- if n < 0 then answer is -1 * n
	- if n > 0 then answer is n
5) Actual code below:
*/

func main() {
	var n float64 = -1

	// 1st way

	// if n < 0 {
	// 	fmt.Println(-1 * n)
	// } else {
	// 	fmt.Println(n)
	// }

	// 2nd way

	// a := math.Float64frombits(math.Float64bits(n) &^ (1 << 63))
	// fmt.Println(a)

	// 3rd way

	// fmt.Println(math.Copysign(n, 1))

	// 4th way

	fmt.Println(math.Abs(n))
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 0, 2, -2
	- Edge cases: Special ones with NaN and Inf
	Tip: We don't have any type conversions implicitly in Go, so using + does not work, we have to use -1 to change the sign.
*/

/*
7) Maintenance:
	- Can I improve the performance? Yes I can. I can use the built-in functions of math package, such as math.Float64Bits() and similar
	  ones to represent the bits of the float input and then change the sign bit to 0, to make it positive. Or I can use the math.Abs()
	  function that is exactly doing the same thing under the hood.
	- Can I make the code cleaner/more readable? Yes, even if I don't want to use the bit representation, I still can use math.Abs() and
	  also I can use math.Copysign(n, 1) to change the sign bit using another built-in function. Either way, we refuse to use branching
	  with if/else statements and we are improving performance and also readability, also making the code shorter.
	- Document anything I change later on for making this code better.
*/

// Takeaway: Understanding what math.Copysign() does, and also understanding two math.Float64bits() and math.Float64frombits() do and
// that we can represent floating-point number's bit representations with them, in terms of IEEE-754 format. And using the 1 << 63 and
// using the &^ (AND NOT) to clear the sign bit of any number is wise!
