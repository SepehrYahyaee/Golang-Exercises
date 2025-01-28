package main

import "fmt"

/*
EX-004: Program to calculate the perimeter and area of a rectangle given it's width and length.
*/

/*
1) Understanding the problem: Simple as before, we're getting a length and width, from a rectangle and we gotta calculate it's area and perimeter.
2) Algorithm:
	inputs: l (float), w (float)
	outputs: p (float), a (float)
	algorithm:
		- Take length (l) and width (w)
		- Calculate perimeter (p) using 2(l+w)
		- Calculate area using l * w
		- Print out 'p' and 'a'
3) Flowchart attached.
4) Pseudo code =>
	- Take l and w
	- p = 2 * (l+w)
	- a = l * w
	- print(l, w)
5) Actual code below:
*/

func main() {
	var l float64 = 0
	var w float64 = 4

	if l <= 0 || w <= 0 {
		fmt.Println("Length/Width could not be negative or zero!")
	} else {
		p := 2 * (l + w)
		a := l * w

		fmt.Println(p, a)
	}
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 2 and 3, 2.1 and 3, 2.5 and 4
	- Edge cases:
		- 0
			+ 0 for either of them is not logical.
		- Negative values
			+ negative values for length/width is not logical.
		- Inf and NaN
			+ as usual
	Tip: Nothing special.
*/

/*
7) Maintenance:
	- Can I improve the performance? Can't think of any way!
	- Can I make the code cleaner/more readable? No.
	- Document anything I change later on for making this code better.
*/
