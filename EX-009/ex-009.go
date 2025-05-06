package main

import "fmt"

/*
EX-009: If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000. (Euler Projects #1)
*/

/*
1) Understanding the problem: Basically, we have to go from 1 to 1000, and find all the multiples of either 3 or 5 and then count the sum of them all. Obviously, we gotta use loops to go from 1 to 1000, and use one form of selection statements (if/else or switch/case/default) to test the numbers, finding out that either they are multiples of 3, or 5.

2) Algorithm:
	inputs: loop's number to count (1000)
	outputs: sum (uint32)
	algorithm:
		- loop through 1 to 999
		- have a variable for summation (0 at first, zero value does it for us)
		- test each input, if it's either a multiple of 3 or 5, add it to the previous sum
		- print the summation at the end of the loop
3) Flowchart attached.
4) Pseudo code =>
	- sum = 0
	- for i = 1; i < 1000; i++ { if i % 3 == 0 or i % 5 == 0 then sum += i, and if not continue on the loop }
	- when the loop ended, print sum
5) Actual code below:
*/

func main() {
	var sum uint32 // Zero value works for us
	var i uint32

	for ; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
		// switch {
		// case i%3 == 0 || i%5 == 0:
		// 	sum += i
		// }
	}

	fmt.Printf("The sum of these multiples is %d\n", sum)
}

/*
6) Testing: it was a calculation process, no testing needed!
*/

/*
7) Maintenance:
	- Can I improve the performance? Yes, the continue command is somewhat unnecessary since the loop automatically goes to the next iteration.
	- Can I make the code cleaner/more readable? blank switch/case style is going to work in the same way too, but the if/else statement appears more readable.
	- Document anything I change later on for making this code better.
*/

// Takeaway: The algorithm was okay to understand, but my assumptions before doing it was that since we are calculating numbers below 1000, the summation is definitely unsigned, and is not going to be more than 500500 ((1000 * 1001) / 2)! So, uint16 can save up to 2 to the power of 16 which is 65535, and is not enough but uint32 is definitely enough. However, if we make the sum variable, typed to uint32, or even uint; since we cannot use declaration form with "var" inside the loops and it has to be with shorthand declaration syntax, so we gotta accept the general default type for untyped literals, which is "int" and we have to accept int64 in here. But what if we declare "i" outside of the loop ? It's actually very well possible, to leave the assignment section of the complete for loop empty and declare it beforehand. Memory-wise, it takes less space to this, although it's not that much but these are all for learning purposes. So this was another performance boost option. (32bits => 4 bytes versus 64bits => 8 bytes => "actually saved 4 bytes :D")
