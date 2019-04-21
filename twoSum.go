package main

import "fmt"

func twoSum(nums []int, target int) []int {

	var sum int
	var s []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]

			if sum == target {
				s = append(s, i)
				s = append(s, j)
				return s
			}
		}
	}

	return s
}

func main() {

	s := []int{3, 2, 4}
	target := 6

	var answer = twoSum(s, target)
	fmt.Println(answer)
}
