package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 14
	step := 0
	//Method 1 correct
	for num != 0 {
		if num%2 == 0 { //偶数
			num = num / 2
			step++
		} else {
			num--
			step++
		}
	}
	fmt.Println("step: ", step)

	num1 := 14
	//Method   -- nor correct
	bNum := strconv.FormatInt(int64(num1), 2)
	fmt.Println("bNum: ", bNum)
	// fmt.Println("left 1:", (bNum >> 1))
	var step2 int = 0
	for num1 != 0 {
		if num1&1 == 1 { //奇数
			num1 = num1 >> 1
			step2 += 2
		} else {
			num1 = num1 >> 1
		}
	}

	fmt.Println("step2: ", step2)
}
