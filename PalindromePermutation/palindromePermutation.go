package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindromePermutation("Tact Coa"))
}

func isPalindromePermutation(str string) bool {

	s := strings.ToLower(strings.Replace(str, " ", "", -1)) //remove empty string and change characters to lowercase

	count := make(map[rune]int)
	//get the count of characters in the string
	for _, c := range s {
		_, ok := count[c]
		if !ok {
			count[c] = 1
		} else {
			count[c]++
		}
	}
	//Check palindrome logic
	even, odd := 0, 0
	for _, v := range count {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if (odd == 0 && even >= 0) || (odd == 1 && even >= 0) {
		return true
	} else {
		return false
	}
}
