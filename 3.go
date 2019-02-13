package main

import "fmt"

func factor(a int) int {

	b := 2
	var c int

	for a != 1 {
		if a%b == 0 {

			a /= b
			c = b
			b = 2
		} else {
			b++
		}
	}
	return c
}

func main() {
	fmt.Println("This is the largest prime factor of 600851475143", factor(600851475143))
}
