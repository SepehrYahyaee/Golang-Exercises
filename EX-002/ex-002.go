package main

import (
	"fmt"
	"math"
)

/*
EX-002: Write a program to print perimeter and area of a circle given it's radius.
*/

/*
1) Understanding the problem: We're given the radius of a circle, and we got to calculate it's area and perimeter.
2) Algorithm:
	inputs: radius (r) - should be float since radius can be decimal.
	outputs: perimeter (p), area (a) - should produce float types.
	algorithm:
		- Take 'r'
		- Calculate 'p' using formula '2 * Pi * r'
		- Calculate 'a' using formula 'Pi * r^2'
		- Print a and p
3) Flowchart attached.
4) Pseudo code =>
	- r (float64)
	- p = 2*Pi*r
	- a = Pi*r*r
	- print(a, p)
5) Actual code below:
*/

func main() {
	var r float64 = /* Change the input as you wish */ 2

	if r < 0 {
		fmt.Println("Radius cannot be negative!")
	} else if math.IsInf(r, 0) {
		fmt.Println("Radius cannot be infinity!")
	} else if math.IsNaN(r) {
		fmt.Println("Radius cannot be NaN!")
	} else {
		p := 2 * math.Pi * r
		a := math.Pi * r * r
		fmt.Println(p, a)
	}
}

/*
6) Testing:
	- Test with some numbers to see the functionality => 0, 2, 5, 7 ...
	- Edge cases:
		- What if the input is negative? we don't have a negative radius around here! (I wonder, in the complex numbers world, is this possible?)
			+ Radius being negative is a real edge case, and I am not going to let that happen by filtering it.
		- What if the input is NaN or positive/negative infinity? that's not logical.
			+ NaN and infinities could be edge cases too but they are some kind of special values and I think letting
			  them be like this is way better, since the user wants to know if something wrong happens, and NaN or a
			  value of infinity appears, but I could prevent that too. It's a matter of preference.
	- So, I gotta prevent negative inputs, which is done simply by an 'if statement', and I want to check for special values too.

	Well, the whole flowchart has been changed due to this conditional statements, but I gotta do it to make sure this app does not produce
	wrong results. I'm not gonna edit the flowchart, that's the whole plan either way, it's just the in the matter of testing stuff, we gotta
	change some things. I could have ignored NaN or infinities too and let them be as special values.

	Tip:
		- Why can't I check a value being NaN, using 'value == math.NaN()' ? It returns false, but using math.IsNaN() is correct! Need to read
		  about that and research. Why NaN does not equal itself?
				+ Reason:
				  Because of specification of IEEE-754 floating point formats, NaN is not going to be equal to itself and the reason is,
				  NaN happens because of mathematical calculations that are not logical in that sense, let's say 0/0 or 0 * Inf or some
				  kind of that stuff, and if a NaN is gonna be equal to any other NaN, it means that 0/0 is equal to 0 * Inf, which is
			      not true, it means that any kind of undefined value no matter how it got created, is equal to some other kind of an
				  undefined value!

				  NaN is for showing some kind of undefined and not logical values, and it is a specification choice to make it not being
				  equal to itself, to prevent considering some awkwardly different expressions being equal to each other, since both of them
				  may produce Not-a-number. So that's why NaN == NaN is false. By definition, NaN is not equal or comparable to itself in any
				  way. math.NaN() returns a NaN and it does not equal itself.

				  Internal representation of NaN as far as I'm aware of, is that the exponential part of it is all 1's and the mantissa part
				  is anything but zero. (Zero would produce Infinities) So any combination of mantissa bits produce NaN, different NaNs! We
				  got two different type of NaNs, qNaNs (quiet NaN) and sNaNs (signaling NaN) which honestly I kinda am not sure what is it
			      that they are different in, but the thing that matters is what math.IsNaN() checks for, is to check a value as in it's
				  internal bit representation to see if it equals any type of NaN, and if it is, it returns true, indication it is a NaN.

				  So, the takeaway is, by definition, NaN is not going to be equal to anything and not equal to itself too, and now I know
				  why. In Go, we cannot directly access NaN, and to do so we gotta use math.NaN() which returns a NaN. And to check if a
				  value is NaN, since they're not comparable, we gotta check the internal bit representation to see if it matches any NaN,
				  and that's exactly what math.IsNaN() does and tells us whether a value is NaN or not. That's why we need to use it this
				  way.
		- '2' is a cool input, both perimeter and area would be the same using this value, cool!
*/

/*
7) Maintenance:
	- Can I improve the performance? not really, I could just reduce the conditional statements by ignoring special values.
	- Can I make the code cleaner/more readable? I don't think so.
	- I need to document anything I change later on for making this code better.
*/

// Takeaway: I understood and researched about why NaN is not equal to itself and how to compare NaNs.
