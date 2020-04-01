package main

import "fmt"

func main() {
	nums := [...]int{8, 1, 2, 2, 5}
	leng := len(nums)
	out := make([]int, leng)
	for i, _ := range nums {
		for j, _ := range nums {
			if nums[i] > nums[j] {
				out[i] += 1
			}
		}
	}
	fmt.Println(out)
}
