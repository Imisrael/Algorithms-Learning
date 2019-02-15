package main

import (
	"fmt"
)

func lcm(x int) int {

	for i := 1; i < 21; i++ {
		if x%i != 0 {
			return 0
		}
	}

	return 1
}

func test(x int) int {

	var smallest int

	for i := x; i > 0; i-- {
		if lcm(i) == 1 {
			smallest = i
		}
	}

	return smallest

}

func main() {

	fmt.Println("man, this really is what it's for: ", test(1000000000))
}
