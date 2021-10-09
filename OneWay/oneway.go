package main

import "fmt"

func main() {
	fmt.Println(oneWay("pale", "ple"))
	fmt.Println(oneWay("pales", "pale"))
	fmt.Println(oneWay("pale", "bale"))
	fmt.Println(oneWay("pale", "bake"))
}

func oneWay(a, b string) bool {
	m := make(map[string]int)

	for _, v := range a {
		if _, ok := m[string(v)]; ok {
			m[string(v)]++
		} else {
			m[string(v)] = 1
		}
	}

	// insert a character
	if len(a) > len(b) {
		for i := 0; i < len(b); i++ {
			if _, ok := m[string(b[i])]; ok {
				m[string(b[i])]--
				if m[string(b[i])] == 0 {
					delete(m, string(b[i]))
				}
			} else {
				continue
			}
		}
		if len(m) == 1 {
			return true
		}
	}

	// remove a character
	if len(a) < len(b) {
		for i := 0; i < len(a); i++ {
			if _, ok := m[string(b[i])]; ok {
				m[string(b[i])]--
				if m[string(b[i])] == 0 {
					delete(m, string(b[i]))
				}
			} else {
				continue
			}
		}
		if len(m) == 0 {
			return true
		}
	}

	// replace a character
	if len(a) == len(b) {
		for i := 0; i < len(b); i++ {
			if _, ok := m[string(b[i])]; ok {
				m[string(b[i])]--
				if m[string(b[i])] == 0 {
					delete(m, string(b[i]))
				}
			} else {
				continue
			}
		}
		if len(m) == 1 {
			return true
		}
	}

	return false
}
