//	The following iterative sequence is defined for the set of positive integers:
//
//	n → n/2 (n is even)
//	n → 3n + 1 (n is odd)
//
//	Using the rule above and starting with 13, we generate the following sequence:
//
//	13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//	It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
//
//	Which starting number, under one million, produces the longest chain?
//
//	NOTE: Once the chain starts the terms are allowed to go above one million.

package main

import "fmt"

func collatzItems(n int) int {

	var total int

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			total++
		} else {
			n = (3*n + 1)
			total++
		}
	}
	return total + 1
}

func main() {

	var answer int

	for i := 999999; i > 1; i-- {
		temp := collatzItems(i)
		if temp > answer {
			answer = i
		}
	}
	fmt.Println("The number with the highest items is: ", answer)

}
