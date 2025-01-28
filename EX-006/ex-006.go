package main

import (
	"fmt"
	"math"
)

/*
EX-006: Program to identify the bigger input between two number inputs.
*/

/*
1) Understanding the problem: We are getting two numbers, and we have to report the maximum between them.
2) Algorithm:
	inputs: a (float), b (float)
	outputs: a or b
	algorithm:
		- Take a and b
		- if a > b => answer is a
		- if b > a => answer is b
		- else => they are equal
3) Flowchart attached.
4) Pseudo code =>
	- a, b
	- if a > b then answer is a
	- if not, if b > a then answer is b
	- if not, they are equal
5) Actual code below:
*/

func main() {
	var a float64 = -0
	var b float64 = -1

	// if a > b {
	// 	fmt.Println(a)
	// } else if b > a {
	// 	fmt.Println(b)
	// } else {
	// 	fmt.Println("equal")
	// }

	// switch {
	// case a > b:
	// 	fmt.Println(a)
	// case b > a:
	// 	fmt.Println(b)
	// case a == b:
	// 	fmt.Println("equal")
	// default:
	// 	fmt.Println("Invalid input: NaN!")
	// }

	fmt.Println(math.Max(a, b))
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 4.5 and 5, -0 and -1
	- Edge cases: Nothing special other than NaN and Inf values, and also precision problems within floats.
	Tip: Nothing special.
*/

/*
7) Maintenance:
	- Can I improve the performance? I don't think so. I don't know if the performance would increase if I change it to the switch/case.
	- Can I make the code cleaner/more readable? Switch/case in my opinion is going to make it better. But overall, the built-in function
	  called math.Max(a, b) is going to suit me more because it's most likely already have the best performance, my previous attemp was
	  to build a new math.Max(a, b) function which is in this point, not probably the right idea to pursue.
	- Document anything I change later on for making this code better.
*/
