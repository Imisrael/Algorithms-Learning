package main

import (
	"fmt"
	"math"
)

func check(a, b, c float64) bool {

	if math.Pow(a, 2)+math.Pow(b, 2) == math.Pow(c, 2) {
		return true
	}

	return false

}

//func pythTriple() float64 {
//
//	var m, n float64
//
//	m = 1
//	n = 2
//
//	for true {
//		a := math.Pow(n, 2) - math.Pow(m, 2)
//		b := 2 * n * m
//		c := math.Pow(n, 2) + math.Pow(m, 2)
//		n++
//		m++
//		if check(a, b, c) {
//			fmt.Println("a, b, c: ", a+b+c)
//			if a+b+c == 1000 {
//				fmt.Println("Did it work?")
//				return 10
//			}
//		}
//	}
//
//	return 10
//}

func backwards(x int) int {

	var result int

	for a := 1; a < x; a++ {
		for b := 2; b < x; b++ {
			c := x - a - b
			fmt.Println("c: total", c, x, a, b, x-a-b)
			if check(float64(a), float64(b), float64(c)) {
				fmt.Println("a, b, c, : ", a, b, c)
				result = a * b * c
				return result
			}
		}
	}

	return result
}

func main() {
	fmt.Println(backwards(1000))
}
