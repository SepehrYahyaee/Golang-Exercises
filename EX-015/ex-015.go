package main

import (
	"fmt"
	"math"
)

/*
EX-015: Write a program to calculate the greatest common divisor of two numbers (gcd).
*/

/*
1) Understanding the problem: Well, gcd of two numbers is the greatest common divisor of them, which is a number definitely smaller than both of them but if we divide each of the targets to that number, the remainder is going to be zero. Since it's the greatest common divisor and not the smallest, it's safe to start from the higher end spiral down.

2) Algorithm:
	inputs: a, b (int)
	outputs: gcd
	algorithm:
		- min = math.Min(a,b)
		- if a%min and b%min == 0 then min is the answer
		- if not min--
		- again go to step 2
		- if min == 0 stop and a,b has no gcd.
3) Flowchart attached.
4) Pseudo code =>
	- min = min(a,b)
	- while min > 0
	- if a%min && b%min == 0 then return min
	- if not then min--
	- go to step 2
5) Actual code below:
*/

func main() {
	a, b := 1540, 1640
	a, b = int(math.Abs(float64(a))), int(math.Abs(float64(b)))

	// First way
	// if a == 0 || b == 0 {
	// 	fmt.Println("a or b cannot be 0")
	// } else {
	// for min := int(math.Min(a, b)); min > 0; min-- {
	// 		if a%min == 0 && b%min == 0 {
	// 			fmt.Println(min)
	// 			break
	// 		}
	// 	}
	// }

	// Second way
	// if a == 0 {
	// 	fmt.Println(a)
	// } else if b == 0 {
	// 	fmt.Println(b)
	// } else {
	// 	for a != b {
	// 		if a > b {
	// 			a -= b
	// 		} else {
	// 			b -= a
	// 		}
	// 	}
	// 	fmt.Println(a)
	// }

	// Third way
	if a == 0 || b == 0 {
		fmt.Println("a or b cannot be 0")
	} else {
		for b != 0 {
			a, b = b, a%b
		}
		fmt.Println(a)
	}
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 91, 52 - -2, -8
	- Edge cases:
		- Negative numbers
			+ For negative numbers, the algorithm would not work well, so we got to use the absolute values everytime. But in the definition, negative numbers isn't going to make a difference in the result, as they're not logical in this context! We're not talking about directions, we are just talking about pure values.
		- 0
			+ having 0 is not logical for finding gcd!
*/

/*
7) Maintenance:
	- Can I improve the performance? Absolutely yes! This algorithm now getting used is the brute force way of calculating gcd. (or hcf - highest common factor) But there are much more better algorithms. Euclidean algorithms are much more efficient. The first algorithm of the Euclidean way is by subtraction. It states that by subtracting two numbers from each other, the gcd stays within them until one of them reaches 0 and then the answer is the greater one! We can also make it more efficient on the second solution too! By using division (remainder) instead of subtraction we can make it even more efficient!
	- Can I make the code cleaner/more readable? I don't think so, it can be written using recursive functions but since we haven't been into functions yet, we won't use that way, and honestly I think it wouldn't be more performant than the loop here (cause of stacks!).
	- Document anything I change later on for making this code better.
*/

// Takeaway: To calculate GCD of two or more numbers, euclidean algorithm of subtraction or furthermore, division; can help us a lot and actually raises the performance to the highest possible state. Trying to prove it also helps us understand why the algorithm works, instead of the first solution.
