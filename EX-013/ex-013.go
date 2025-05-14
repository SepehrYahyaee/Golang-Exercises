package main

import (
	"fmt"
)

/*
EX-013: Write a program to find if a number is palindrome or not.
*/

/*
1) Understanding the problem: A palindromic number reads the same both ways, so we need to give a target number and identify if it reads the same both ways or not and print true or false.
2) Algorithm:
	inputs: target
	outputs: boolean (is palindrome or not)
	algorithm:
		- Modulus of a number with 10 gives us the last digit
		- Divide an integer with 10 and replacing it with itself, removes the last digit
		- reversed = 0
		- While the number is not 0, 10 * reversed = reversed*10 + lastDigit
		- Check if the target is equal to the created reversed
3) Flowchart attached.
4) Pseudo code =>
	- target = num
	- temp = target
	- reversed = 0
	- for temp != 0 { reversed = reversed*10 + temp%10; temp /= 10 }
	- if target == reversed then number is palindrome
5) Actual code below:
*/

func main() {
	var target int = 132465798978645312
	var temp int = target
	var reversed int = 0

	if target < 0 || (target != 0 && target%10 == 0) {
		fmt.Print("Target is not palindrome.")
	} else {
		for temp > reversed {
			reversed = reversed*10 + temp%10
			temp /= 10
		}

		if temp == reversed || temp == reversed/10 {
			fmt.Print("Target is palindrome.\n")
		} else {
			fmt.Print("Target is not palindrome.\n")
		}
	}
}

/*
6) Testing:
	- Test with some inputs to see the functionality =>
	- Edge cases:
		- Negative numbers!
			+ Negative numbers like -1221 gives us the palindrome result but actually they're not! Palindrome numbers are not negative at all.
		- NaN and Infinity and special stuff actually wouldn't work (they're float64 values)
*/

/*
7) Maintenance:
	- Can I improve the performance? I guess so; Numbers that have 0 at their end, in other words are fully dividable by 10 cannot be palindrome, since the very first digit of a number cannot be 0! So we could check those out in the first if statement already, instead of going with the flow. But remember, 0 is a palindrome itself so you gotta separate that out. We could also cut the loop in half, by just reversing the second half for the equality.
	- Can I make the code cleaner/more readable? The code is pretty readable but since we haven't studied functions; it'll be looking like this for a while.
	- Document anything I change later on for making this code better.
*/

// Takeaway: Modulus%10 and you would have the last digit; Target /= 10 and you would delete the last digit effectively; and to build a number from start you have to multiply the digit in 10 and then add the digit you want. (Mathematics stuff) Also temp == reversed || temp == reversed/10 is a condition checking both even and odd numbers being palindrome; evens are temp == reversed and odds are temp == reversed/10 [middle digit does not matter] We can do it with strings and stuff too which we'll do later on, but they are not going to be performant enough!
