package main

import "fmt"

func makeRange(min, max int) []int { // found this online

	x := make([]int, max-min+1)
	for i := range x {
		x[i] = min + i
	}
	return x
}

func sumOfSquares(x []int) int {

	n := len(x)
	var total int

	for i := 0; i < n; i++ {
		total = x[i]*x[i] + total
	}
	return total
}

func squareOfSum(x []int) int {

	n := len(x)
	var total int

	for i := 0; i < n; i++ {
		total += x[i]
	}
	total = total * total
	return total
}

func main() {

	x := makeRange(1, 100)

	fmt.Println("the difference between the two numbers: ", squareOfSum(x)-sumOfSquares(x))

}
