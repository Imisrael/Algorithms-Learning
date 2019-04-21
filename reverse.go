package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(n int) int {

	var s []string

	abso := math.Abs(float64(n))
	x := int(abso)

	for x > 0 {
		temp := x % 10
		y := strconv.Itoa(temp)
		s = append(s, y)
		x = x / 10
	}

	temp := strings.Join(s, "")
	num, _ := strconv.Atoi(temp)

	if n < 0 {
		num = num * -1
	}

	if num > 2147483647 || num < -2147483647 {
		return 0
	}

	return num
}

func main() {
	fmt.Println(reverse(-2147483648))
}
