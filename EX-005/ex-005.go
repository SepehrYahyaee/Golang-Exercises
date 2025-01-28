package main

import "fmt"

/*
EX-005: Program to calculate the area of a triangle given it's height and base.
*/

/*
1) Understanding the problem: We are getting a height and a base for a triangle, and want to calculate it's area.
2) Algorithm:
	inputs: h (float), b (float)
	outputs: a (float)
	algorithm:
		- Take height (h) and base (b)
		- Calculate area (a) using a = 1/2(h + b)
		- Print out a
3) Flowchart attached.
4) Pseudo code =>
	- h, b
	- a = 1/2*(h + b)
	- print(a)
5) Actual code below:
*/

func main() {
	var h float64 = 3
	var b float64 = 2

	if h <= 0 || b <= 0 {
		fmt.Println("Height/Base cannot be negative or zero!")
	} else {
		a := 0.5 * (h + b)
		fmt.Println(a)
	}
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 2 and 3, 3 and 2
	- Edge cases:
		- 0
			+ 0 for lengths?
		- Negative values
			+ Negative values are not for lengths.
		- NaN or Inf
			+ as usual, treating them as special values.
	Tip:
		- Using 0.5, instead of 1/2
			+ for the same reason as in ex-003, using less instructions, and also not using untyped integers that would ruin the
			  calculations.
*/

/*
7) Maintenance:
	- Can I improve the performance? Yes, changing 1.0/2.0 to 0.5, also I can literally not make a variable and just print the
	  calculated value directly inside the fmt.Println() function, but let that be for more readability.
	- Can I make the code cleaner/more readable? No.
	- Document anything I change later on for making this code better.
*/
