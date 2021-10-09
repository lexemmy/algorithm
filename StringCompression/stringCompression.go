package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(compression("aabcccccaaa"))
}

func compression(s string) string {
	count := 1
	str := ""
	for i := 0; i < len(s)-1; i++ {
		if string(s[i]) == string(s[i+1]) {
			count++
		} else {
			c := strconv.Itoa(count)
			str += string(s[i]) + c
			count = 1
		}
	}
	c := strconv.Itoa(count)
	str += string(s[len(s)-1]) + c

	if len(str) > len(s) { //if compressed string is not smaller than original string, return original string
		return s
	}
	return str
}
