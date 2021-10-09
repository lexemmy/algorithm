package main

import "fmt"

func main() {
	fmt.Println(isUnique("youverify"))
	fmt.Println(isUnique2("identity"))
}

//without using additional data structure, the solution is not effiecient and the complexity is O(n^2)
func isUnique(str string) bool {
	for i := 0; i < len(str); i++ {					//loop through the string
		for j := i + 1; j < len(str); j++ {			//a second loop through the string
			if str[i] == str[j] {					//check if a character is equal to another charcter in the string 
				return false						//return false if two characters are equal
			}
		}
	}
	return true
}

//with data additional structure, the solution is efffiecient and the complexity is O(n)
func isUnique2(str string) bool {
	m := make(map[string]int)						//make an empty map
	for i := 0; i < len(str); i++ {					//loop through string
		if _, ok := m[string(str[i])]; ok {			//check if a character exist in map 
			return false							//return false is character already exist
		} else {
			m[string(str[i])]++						//strore the character in the map if it does not exist
		}
	}
	return true
}
