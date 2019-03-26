package main

import (
	"fmt"
	"math"
)

func numOfDivisors(x int) float64 {

	var numOfFactors float64
	var sqrt = math.Sqrt(float64(x))

	for i := 1; float64(i) <= sqrt; i++ {
		if x%i == 0 {
			if float64(i) == sqrt {
				numOfFactors++
			} else {
				numOfFactors += 2
			}
		}
	}
	return numOfFactors
}

func triangleNum(x int) int {

	x = x * (x + 1)
	x /= 2

	return x

}

func main() {

	i := 1

	for {
		temp := triangleNum(i)
		temp2 := numOfDivisors(temp)
		fmt.Println("temp: ", temp, "temp2: ", temp2)
		if temp2 > 500 {
			fmt.Println(temp)
			return
		}
		i++
	}
}
