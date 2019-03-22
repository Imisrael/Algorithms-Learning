package main

import (
	"fmt"
)

func factor(x int) []int {

	var result []int

	for i := 1; i < x+1; i++ {
		if x%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func triangleNum(x int) int {

	var total int

	for i := 1; i < x+1; i++ {
		total += i
	}

	return total

}

func main() {

	i := 1

	for {
		temp := triangleNum(i)
		temp2 := factor(temp)
		if len(temp2) > 500 {
			fmt.Println(temp)
			return
		}
		i++
	}
}
