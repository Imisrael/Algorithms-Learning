package main

import (
	"fmt"
)

func main() {

	fmt.Println(fibonacci())
}

func fibonacci() int {
	a := 1
	b := 2
	var c int
	var total int

	for b <= 4000000 {
		if b%2 == 0 {
			fmt.Println("b: ", b)
			total += b
		}
		c = a + b
		a = b
		b = c
	}
	return total

}
