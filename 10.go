package main

import (
	"fmt"
)

//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
//Find the sum of all the primes below two million.

func isPrime(x int) bool {

	if x <= 1 {
		return false
	}

	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func iteratePrime(x int) int {

	var total int

	for i := 0; i < x; i++ {
		if isPrime(i) {
			total += i
			fmt.Println("total: ", total, i)
		}
	}
	return total
}

func main() {

	fmt.Println("This is the answer: ", iteratePrime(2000000))
}
