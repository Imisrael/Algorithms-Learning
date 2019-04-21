package main

import (
	"fmt"
)

func checkPossibility(nums []int) bool {

	n := len(nums)
	var count int

	if n == 1 {
		fmt.Println("n == 1")
		return true
	}

	for i := 0; i < n-1; i++ {

		if nums[i] > nums[i+1] || nums[i] > nums[i-1] {
			for j := 1; j < 9; j++ {
				if j <= nums[i+1] {
					fmt.Println("j: ", j)
					nums[i] = j
					count = 0
					for k := 0; k < n-1; k++ {
						if nums[k] <= nums[k+1] {
							count++
							fmt.Println("count: ", count, k, k+1)
						}
						if count == len(nums)-1 {
							fmt.Println("sub loop")
							return true
						}
					}
					fmt.Println("ending sub loop", nums)
					return false
				}
			}
		}

		if nums[i] <= nums[i+1] {
			count++
		}
		if count == len(nums)-1 {
			fmt.Println("got to the end")
			return true
		}
	}

	fmt.Println("very ending")
	return false
}

func main() {

	s := []int{2, 3, 3, 2, 4}
	fmt.Println(checkPossibility(s))

}
