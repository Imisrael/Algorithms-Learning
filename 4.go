package main

import "fmt"

func palindrome(x int) bool {

	temp := x
	reverse := 0
	var dig int

	for x > 1 {
		dig = x % 10
		reverse = reverse*10 + dig
		x /= 10
	}
	if temp == reverse {
		return true
	} else {
		return false
	}

}

func largest() {

	var largest int

	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			temp := j * i
			if palindrome(temp) == true {
				if temp > largest {
					largest = temp
				}
			}
		}
	}

	fmt.Println("the largest is: ", largest)
}

func main() {

	largest()
}
