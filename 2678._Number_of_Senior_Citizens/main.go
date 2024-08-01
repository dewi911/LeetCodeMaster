package main

import "fmt"

func main() {
	var details []string
	details = append(details, "7868190130M7522", "5303914400F9211", "9273338290F4010")

	result := countSeniors(details)
	fmt.Println(result)
}

func countSeniors(details []string) int {
	if len(details) == 0 {
		return 0
	}

	result := 0
	for i := 0; i < len(details); i++ {
		if int(details[i][len(details[i])-4]-'0')*10+int(details[i][len(details[i])-3]-'0') > 60 {
			if string(details[i][len(details[i])-5]) == "F" || string(details[i][len(details[i])-5]) == "M" || string(details[i][len(details[i])-5]) == "O" {
				result++
			}
		}
	}

	return result
}
