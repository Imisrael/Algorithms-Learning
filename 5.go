package main

import (
	"fmt"
)

func loop(x int) int {

	for i := 1; i < 21; i++ {
		if x%i != 0 {
			return 0
		}
	}

	return 1
}

func lcm(x int) int {

	var smallest int

	for i := 21; i < x; i++ {
		if loop(i) == 1 {
			smallest = i
			break
		}
	}

	return smallest

}

func main() {

	fmt.Println("Smallest Multiple: ", lcm(1000000000))
}
