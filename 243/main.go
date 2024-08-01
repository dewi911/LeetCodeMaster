package main

import (
	"fmt"
)

func main() {
	var temp []string

	temp = append(temp, "boice", "bob", "boarlie")
	fmt.Println(temp)
	govn := longestCommonPrefix(temp)
	fmt.Println(govn)
}

func longestCommonPrefix(s []string) string {
	minLen := len(s[0])
	for i := 0; i < len(s); i++ {
		if len(s[i]) < minLen {
			minLen = len(s[i])
		}
	}

	result := ""
	for i := 0; i < minLen; i++ {
		state := true
		for j := 1; j < len(s); j++ {
			if s[j][i] != s[j-1][i] {
				state = false
			}
		}
		if state {
			result += string(s[0][i])
		} else {
			break
		}
	}
	return result
}
