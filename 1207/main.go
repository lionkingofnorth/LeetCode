package main

import "fmt"

func main() {
	var arr = [...]int{1, 2}
	// var manyTime bool
	// manyTime = true
	//main question
	var countM map[int]int
	countM = make(map[int]int)
	for _, v := range arr {
		if _, ok := countM[v]; !ok {
			countM[v] = 1
		} else {
			countM[v]++
		}
	}
	var occurence []int
	// var s string
	for _, value := range countM {
		occurence = append(occurence, value)
		// s += strconv.Itoa(value) + ","
	}

	fmt.Println(occurence)
	for i := 0; i < len(occurence); i++ {
		for j := i + 1; j <= len(occurence)-1; j++ {
			if occurence[i] == occurence[j] {
				fmt.Println("false")
			}
		}
	}

	//second method
	/*
		numbers := make(map[int]int)
		ocs := make(map[int]bool)
		for _, v := range arr {
			numbers[v]++
		}
		for _, v := range numbers {
			if _, ok := ocs[v]; !ok {
				ocs[v] = true
			} else {
				return false
			}
		}
		return true
	*/

}
