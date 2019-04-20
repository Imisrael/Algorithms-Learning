// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

// What is the sum of the digits of the number 2^1000?

package main

import (
	"fmt"
	"math"
)

func reverseSlice(s []float64) []float64 {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func theSum(num float64) float64 {

	var digits []float64
	var sum float64

	y := math.Pow(2, num)

	for y > 1 {
		temp := math.Mod(y, 10)
		digits = append(digits, temp)
		//fmt.Printf("y: %f\n", y)
		y = y / 10
	}

	digits = reverseSlice(digits)

	for i := 0; i < len(digits); i++ {
		sum += digits[i]
	}

	return sum
}

func main() {

	fmt.Println(theSum(1000))

}
