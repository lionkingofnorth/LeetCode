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
	decode = strings.ReplaceAll(ip, ".", "[.]")
	fmt.Println("ip: ", ip)
	fmt.Println("decode: ", decode)
}
