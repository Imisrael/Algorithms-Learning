package main

import "fmt"

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

	var counter int

	for i := 0; i < 100; i++ {
		if isPrime(i) {
			counter++
			if counter == x {
				return i
			}
		}
	}
	return counter
}

func main() {

	fmt.Println("10,001: ", iteratePrime(10001))
}
