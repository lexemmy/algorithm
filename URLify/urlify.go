package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(urlify("Mr John Smith "))
}

func urlify(str string) string {
	trimedString := strings.Trim(str, " ")
	newString := strings.Replace(trimedString, " ", "%20", -1) //call replace method to replace empty space with %20 for all occurence
	return newString
}
