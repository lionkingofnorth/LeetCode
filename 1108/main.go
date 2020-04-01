package main

import (
	"fmt"
	"strings"
)

var ip string
var decode string

func main() {
	ip = "1.1.1.1"
	// for i, v := range ip {
	// 	fmt.Printf("i: %d v:%T ", i, v)
	// }
	//Method 1
	decode = strings.ReplaceAll(ip, ".", "[.]")
	fmt.Println("ip: ", ip)
	fmt.Println("decode: ", decode)
	//Method 2
	var decode2 string
	for _, v := range ip {
		if string(v) == "." {
			decode2 += string("[.]")
		} else {
			decode2 += string(v)
		}
	}
	fmt.Println("decode2: ", decode2)

}
