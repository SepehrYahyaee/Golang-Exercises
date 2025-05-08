package main

import (
	"fmt"
	"math"
)

/*
EX-011: Program to find out that a given number is prime or not.
*/

/*
1) Understanding the problem: Prime numbers are the integers that are divisible by only themselves and 1, and no other numbers. So, we need to write a program that'll calculate all the prime numbers up to a given point. We got to loop from 2 to the given number and check all the numbers between that the number is divisible by that, or not; if the number was divisible by even 1 number, the number is not prime and if not, that number's prime.

2) Algorithm:
	inputs: target (uint)
	outputs: true/false
	algorithm:
		- Loop from 2 to the target
		- Test if the number is divisible by that number
		- If not divisible by all the numbers, then the number is prime.
		- If divisible by at least one number in between, the number is not prime.
3) Flowchart attached.
4) Pseudo code =>
	- target = number we want to check if prime
	- for i := 2; i < target ; i++ { if target%i == 0 false and break; else continue }
	- if loop ended on it's own, then true
5) Actual code below:
*/

func main() {
	var target int = 9_823_894_127_391_823

	if target == 1 || target < 1 {
		fmt.Println("Target should be positive and above 1")
	}

	// First way

	// for i := 2; i < target; i++ {
	// 	if target%i == 0 {
	// 		fmt.Println(false)
	// 		break
	// 	} else {
	// 		if i == target-1 {
	// 			fmt.Println(true)
	// 		}
	// 	}
	// }

	// Second way

	var isPrime bool = true

	for i := 2; i < int(math.Sqrt(float64(target))); i++ {
		if target%i == 0 {
			isPrime = false
			break
		}
	}

	fmt.Println(isPrime)
}

/*
6) Testing:
	- Test with some inputs to see the functionality => 5, 3, 4, 255, 271, 177
	- Edge cases:
		- 0, 1 (and 2, using the first way to solve)
			+ actually the code would do nothing but should report an error. 1 is also prime anyway.
		- Negative numbers
			+ same as 0.

	Tip:
		- Since we're using the main function and it can't have any return statements, we have to write the program this way, to see if the loop has ended.
			+ We also can use "flag-based" programming which till the end of the loop, flag would be responsible as the result of the calculations, true/false.
		- When we'll learn functions later on, we could easily write the solution with return statements.
*/

/*
7) Maintenance:
	- Can I improve the performance? Yes. First thing is, since finding a factor is enough for a number to not being prime, so using "break" after flag = false is going to omit the next iterations which is useless. Second, there's a mathematical thing about finding primes. "We can safely check only up to the square root of numbers divisibility." What does that mean ? It's being said that you can calculate the square root of a number, and make it a pin point; then if a the target is divisible by a number below that pinpoint, then surely it's divisible by a number above that pinpoint too, and vice versa. So we only need to check numbers up to the square root and can be safely if the number is already prime, then the remaining numbers cannot change the result. Cool stuff to be honest.
	- Can I make the code cleaner/more readable? It's already clean in my opinion. The only thing is, instead of having a "flag" boolean value, renaming it to "isPrime" is much more readable and understandable. Why don't we make isPrime, "false" at the beginning? Because we have to assume the number is prime to find out if one factor is found, change it to false; if it was false, we had to check all the options to find out it's prime, which is not a good way to do it.
	- Document anything I change later on for making this code better.
*/

// Takeaway: The concept of the square root of a number as a pinpoint to find out if a number is divisible by any number greater than that pinpoint is awesome. And that target I wrote up there, is a fucking random input I just typed out of nowhere to see the performance of this code; and holy shit that was prime out of nowhere, I couldn't believe so I checked it up, and yeah it's a prime, damn. Mathematics is fun.
