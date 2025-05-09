package main

import (
	"fmt"
)

/*
EX-012: The prime factors of 13195 are 5, 7, 13 and 29. What is the largest prime factor of the number 600851475143? (Euler Projects #3)
*/

/*
1) Understanding the problem: As we've already solved an exercise on the prime numbers, so we're familiar with it. Now, the prime factors of a number is the numbers which the target number is divisible to, but those are primes. For example, prime factors of number 10, is 2 and 5. Now the question wants us to calculate the largest prime factor of the target number.

2) Algorithm:
	inputs: target (uint)
	outputs: largest prime factor (uint)
	algorithm:
		- loop through 2 to target
		- find primes
		- if target is divisible to that prime, so we save it as largest
		- we'll do this till we find the largest one
3) Flowchart attached. (Only first way)
4) Pseudo code =>
	- Loop from 3 to target (including itself)
	- Check if the number's prime then check if it's divisible by target, if yes then save it as the current largest prime factor
	- To check the primality, we need another loop from 2 to the first loops index.
	- We can do it vice versa (Divisibility, then Primality which is the second way)
5) Actual code below:
*/

func main() {
	var target = 600851475143
	var largestPrimeFactor int

	// First way - Primality then Divisibility

	// for i := 3; i <= target; i++ {
	// 	var isPrime bool = true

	// 	for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
	// 		if i%j == 0 {
	// 			isPrime = false
	// 			break
	// 		}
	// 	}

	// 	if isPrime == true {
	// 		if target%i == 0 {
	// 			largestPrimeFactor = i
	// 		}
	// 	}
	// }

	// fmt.Println(largestPrimeFactor)

	// Second way - Divisibility then Primality

	// for i := 2; i <= target; i++ {
	// 	var isPrime bool = true

	// 	if target%i == 0 {
	// 		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
	// 			if i%j == 0 {
	// 				isPrime = false
	// 				break
	// 			}
	// 		}
	// 		if isPrime == true {
	// 			largestPrimeFactor = i
	// 		}
	// 	}
	// }

	// fmt.Println(largestPrimeFactor)

	// Third way - Factorization
	for i := 2; target > 1; i++ {
		for target%i == 0 {
			target /= i
			largestPrimeFactor = i
		}
	}

	fmt.Println(largestPrimeFactor)
}

/*
6) Testing: it was a calculation process, no testing needed!
*/

/*
7) Maintenance:
	- Can I improve the performance? Yes, we are now checking the divisibility after the primality. If we first check the divisibility and then the primality, it would be much less instructions, and is more logically faster. (Second way) But the whole approach is flawed! First way and second way woudln't work properly and using int(math.Sqrt(float64(target))) in the outer loop is also wrong, because the largest prime factor might be larger than that pinpoint we were talking about in the last exercise! However, magically, out of coincidence, those two ways give us the right answer because we got lucky. It's not always that we get lucky and when we would want to try other numbers, we might not get that lucky (try 91). So the outer loop's condition has to be i <= target, and this way, if the number's large, like the target the question wants from us to calculate; it's not going to work and it's gonna take way more than usual. It's time to learn more about primes, mathematically then to find the best algorithm. (Good luck trying to find the answer using those methods!)

	Mathematically, each number has building blocks of prime numbers; take 10 for example: If we start from the first prime (2) and remove all the factors of 2 from the target number (10), there remains 5; then if we move to next numbers, 3 in here, there's no 3 factor in 5; then 4, yet none, and then 5 itself which makes the target 1 and so the largest prime factor of 10 is 5! It's like peeling the prime factors of any number bit by bit to find out the largest prime factor! Why not checking primes, and checking all the numbers (including 4 in here)? Because we are slowly removing all the prime factors, and by removing 2's factor; automatically all the other numbers like 4,6,8, ... would be removed from the number; so we don't even have to check for primality! Then we do this by moving towards the number and reducing the target every time, and what remains, must be a prime and is the largest factor of it too. (Third way) Take a look at it; mathematics is really fascinating!
	- Can I make the code cleaner/more readable? I don't think I can make it more readable, the third way is awesome.
	- Document anything I change later on for making this code better.
*/

// Takeaway: Brute forcing your way into the problem might solve the question for the low inputs, but those are not the ultimate answer; we must strive for the best algorithm at all times. I learnt that numbers in mathematics has been made up of prime numbers that if we slowly peel those primes out of a number, then the largest prime factor of a number remains, cool stuff!
