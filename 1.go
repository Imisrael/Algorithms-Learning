package main

import "fmt"

func main() {

	fmt.Println(findMulti(1000))
}

func findMulti(n int) int {
	var total int
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total

}
