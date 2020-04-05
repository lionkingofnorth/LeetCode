package main

import (
	"fmt"
	"strconv"
)

/*
Given an integer number n, return the difference between the product of its digits and the sum of its digits.


Example 1:

Input: n = 234
Output: 15
Explanation:
Product of digits = 2 * 3 * 4 = 24
Sum of digits = 2 + 3 + 4 = 9
Result = 24 - 9 = 15
Example 2:

Input: n = 4421
Output: 21
Explanation:
Product of digits = 4 * 4 * 2 * 1 = 32
Sum of digits = 4 + 4 + 2 + 1 = 11
Result = 32 - 11 = 21

Constraints:

1 <= n <= 10^5

*/

func main() {
	n := 234
	// nr := []rune(n)
	// fmt.Println(nr)
	data := strconv.Itoa(n)
	var product, sum int
	product = 1
	sum = 0
	for i := range data {
		fmt.Println(data[i])
		temp, _ := strconv.Atoi(string(data[i]))
		product *= temp
		sum += temp
	}
	fmt.Println("Product is : ", product)

	fmt.Println("Sum is : ", sum)
	fmt.Println("Sub is : ", product-sum)
}
