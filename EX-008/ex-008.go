package main

import (
	"fmt"
)

/*
EX-008: Program to calculate total resistance of three resistors in an electrical circuit using the given formula.
*/

/*
1) Understanding the problem: We just have a formula and three input numbers as resistance.
2) Algorithm:
	inputs: R1, R2, R3 (float)
	outputs: T (float)
	algorithm:
		- Take R1, R2 and R3
		- Calculate T using the formula
		- Proint out T
3) Flowchart attached.
4) Pseudo code =>
	- R1, R2, R3
	- T = 1 / (1/R1 + 1/R2 + 1/R3)
	- Print(T)
5) Actual code below:
*/

func main() {
	var R1, R2, R3 float64 = 1.5, 2.5, 0

	if R1 <= 0 || R2 <= 0 || R3 <= 0 {
		fmt.Println("Negative/Zero resistance is not allowed!")
	} else {
		// T := 1 / ((1 / R1) + (1 / R2) + (1 / R3))

		fmt.Println(R1 * R2 * R3 / (R2*R3 + R1*R3 + R1*R2))
	}
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 1.5 and 2.5 and 3.5
	- Edge cases:
		- Negative values and Zero.
		- NaN and Inf, which I ignore just like the rest of the problems.
	Tip: It's interesting that n/Inf = + or - zero
*/

/*
7) Maintenance:
	- Can I improve the performance? Yes, by transforming the formula mathematically, it's most likely that the performance is being
	  increased since division costs more than multiplication, and in this case, I decreased 4 divisions to one, and added 5 instruction
	  of multiplication. I also can just ignore the created variable and directly print the value.
	- Can I make the code cleaner/more readable? I don't think so.
	- Document anything I change later on for making this code better.
*/

// Takeaway: It's not that big of a deal, but division costs slightly more than multiplication in terms of CPU instructions, so if you're
// able to use anything else instead of division operators, doing it would benefit the code and the system.
