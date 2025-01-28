package main

import (
	"fmt"
)

/*
EX-003: Program to convert a given centigrade temperature, to it's fahrenheit equivalent.
*/

/*
1) Understanding the problem: We are getting a temperature number in celsius units, and it needs to be converted to fahrenheit.
2) Algorithm:
	inputs: c (float)
	outputs: f (float)
	algorithm:
		- Take C (celsius)
		- Convert it to fahrenheit using the formula F = 9/5C + 32
		- Print out F
3) Flowchart attached.
4) Pseudo code =>
	- C
	- F = 9/5*C + 32
	- Print(F)
5) Actual code below:
*/

func main() {
	var c float64 = /* Or any input you want */ 0
	f := 1.8*c + 32
	fmt.Println(f)
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 0, 2.5, -10, math.MaxFloat64
	- Edge cases:
		- NaN
			+ returns NaN for any input
		- Positive/Negative infinity
			+ return +/- Inf

	Tip:
		- If we use literals like 9/5 casually, since they are untyped, the result would be 1*c which is wrong!
			+ Since the result is float and we want to make sure the result is correct, we have to use 9.0/5.0 to make sure the division
			  is correct and returns a correct value.
		- math.SmallestNonzeroFloat64
			+ If I use this value, the result will be calculated as if the input was 0, If we +1 the input, it is actually 1! Now if we
			  print the math.SmallestNonzeroFloat64 value as itself, it will print out 5e-324, But in the calculations it will work out
			  as 0, like the compiler looks at it as zero, just like it looks at the big numbers as Inf.
*/

/*
7) Maintenance:
	- Can I improve the performance? Yes, just use 1.8 for 9/5 and you make one less instruction to calculate, also you don't have the untyped
	  constants problem anymore, since 1.8 is already float.
	- Can I make the code cleaner/more readable? Nah.
	- Document anything I change later on for making this code better.
*/

// Takeaway: math.SmallestNonzeroFloat32/64 actually prints out their value, but using them in calculations is the same az zero! Just like we
// have Overflow, we have Underflow too and that is, collapsing to 0, and in case of overflow, the behaviour is saturation to Inf/-Inf. The
// takeaway here is that I understood Overflow/Underflow in IEEE-754 representations and it's because of it's limited bits.
