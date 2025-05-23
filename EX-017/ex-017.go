package main

import (
	"fmt"
	"math"
)

/*
EX-017: 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder. What is the smallest positive number that is evenly divisible with no remainder by all of the numbers from 1 to 20? (Euler Projects #5)
*/

/*
1) Understanding the problem: Well, the exercise asks us to calculate a number that is divisible by all the numbers from 1 to 20, must be a common factor to them all, and also the exercise wants the smallest of them, meaning that we gotta calculate the LCM of 1 to 20. The thing is, LCM(1, ..., 20) can be written like LCM(LCM(1,2), 3 ...) mathematically. Meaning that if we calculate the LCM of all the numbers starting from 1 and 2, and then their result with 3 and so on, we reach the LCM of all the numbers of 1 to 20.

2) Algorithm:
	inputs: -
	outputs: LCM(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	algorithm:
		- for i := 1; i <= 20; i++ { calculate LCM of i and i+1 }
		- Print the result at the end
3) Flowchart attached.
4) Pseudo code =>
	- lcm = 1
	- for i := 2; i <= 20; i++ { a, b := lcm, i }
	- calculating GCD
	- lcm = lcm*i / a
	- print lcm at the end
5) Actual code below:
*/

func main() {

	// First way
	// lcm := 1

	// for i := 2; i <= 20; i++ {
	// 	a, b := lcm, i

	// 	for b != 0 {
	// 		a, b = b, a%b
	// 	}

	// 	lcm = lcm * i / a
	// }

	// fmt.Println(lcm)

	// Second way
	result := 1

	for i := 2; i <= 20; i++ {
		isPrime := true

		for j := 2; j < i; j++ {
			if i%j == 0 && i != j {
				isPrime = false
				break
			}
		}

		if isPrime == true {
			power := int(math.Floor(math.Log(float64(20)) / math.Log(float64(i))))
			result *= int(math.Pow(float64(i), float64(power)))
		}
	}

	fmt.Println(result)
}

/*
6) Testing: it was a calculation process, no testing needed!
*/

/*
7) Maintenance:
	- Can I improve the performance? I don't think so, for the first solution. But there's another solution that might be faster than the first one for the bigger numbers or as a whole, and it's worth knowing the algorithm. Number theory says, every number above 1 is either prime or can be made from the products of primes (obviously primes below itself). Now, to find the least common multiple of numbers from 1 to let's say k, we know at most that the result could be 1*2*3*4*...*20 but these numbers have many common factors with each other that we'll have to divide or not count, so that we'll find the real LCM. Primes are different in a series of numbers, from 1 to k. Those numbers don't have common stuff with others so we can safely count them and multiply them, but for other numbers that aren't prime, we only need to divide them to their prime factors and then add the prime ones to the previous existing primes. In other words, if we find the highest available power of a prime in a series of numbers, and only count primes in that series and multiply all of them together, the result is the LCM of them all! But we gotta look for overflows in the second solution since it grows fast and we might end up with negative results which is not correct. We got to understand overflows of the algorithm and prepare for them. For example, LCM of all numbers from 1 to 100 might not even fit into the 64bits, even unsigned. So we gotta work with math/big library which is another story.
	- Can I make the code cleaner/more readable? Yeah if we learn functions, but for now, it's all good.
	- Document anything I change later on for making this code better.
*/

// Takeaway: LCM of more than two numbers can be divided by LCM of all of them two by two, together. Mathematics allow that. Imagine if you wanted to calculate such a thing using brute force algorithms for LCM :). Also, prime numbers and number theory can help a lot as you can see in the second solution, they are fun to learn.
