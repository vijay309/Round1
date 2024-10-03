package main

import (
	"fmt"
)

func main() {
	var s string
	var k int

	// Input the string and rotation count
	fmt.Print("Enter the string: ")
	fmt.Scan(&s)
	fmt.Print("Enter the rotation count: ")
	fmt.Scan(&k)

	// Adjust k if it exceeds the length of the string
	if k > len(s) {
		k = k - len(s)
	}

	lenS := len(s)
	size := lenS - k
	res := ""
	for i := k; i < lenS; i++ {
		res += string(s[i])
	}
	for j := 0; j < k; j++ {
		res += string(s[j])
	}
	
	if k == len(s) {
		res = s
	}
	fmt.Println("Final result:", res)
} 