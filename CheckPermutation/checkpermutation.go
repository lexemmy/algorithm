package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(checkPermutation("verify","eriyfv"))
	fmt.Println(checkPermutation("you","me"))
}

//a function to convert string to rune
func stringsToRunes(str string) []rune {
	r := []rune{}
	for _, v := range str {
			r = append(r, rune(v))
	}
	return r
}

func checkPermutation(str1,str2 string) bool {
	if len(str1) != len(str2) { 						//check if the length of the string are equal and return false if true
		return false
	}
	newStr1 := stringsToRunes(str1)						//convert the string to rune
	newStr2 := stringsToRunes(str2)
	sort.SliceStable(newStr1, func(i, j int) bool {		//sort the rune 
		 return newStr1[i] < newStr1[j]
	})
	sort.SliceStable(newStr2, func(i, j int) bool {
		 return newStr2[i] < newStr2[j]
	})
	if string(newStr1) != string(newStr2) {				//convert rune back to string and check if equal
		return false
	}
	return true
}

