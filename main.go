package main

import (
	"fmt"
	"strings"
)

func possibilities(signals string) []string {
	var result []string

	// Create a map to store the Morse code for each letter
	morseCode := map[string]string{
		".":   "E",
		"..":  "I",
		"...": "S",
		".-":  "A",
		"..-": "U",
		".-.": "R",
		".--": "W",
		"-":   "T",
		"-.":  "N",
		"-..": "D",
		"-.-": "K",
		"--.": "G",
		"---": "O",
		"--":  "M",
	}
	result = append(result, signals)
	for {
		var contains bool
		for _, s := range result {
			contains = strings.Contains(s, "?")
			if contains {
				withDot := strings.Replace(s, "?", ".", 1)
				withDash := strings.Replace(s, "?", "-", 1)
				result = append(result, withDot)
				result = append(result, withDash)
				result = result[1:]
			} else {
				break
			}
		}
		if !contains {
			break
		}
	}
	for i, s := range result {
		result[i] = morseCode[s]
	}
	return result
}

func main() {
	fmt.Println(possibilities("."))
	fmt.Println(possibilities("-"))
	fmt.Println(possibilities("-."))
	fmt.Println(possibilities("..."))
	fmt.Println(possibilities("..-"))
	fmt.Println(possibilities("?"))
	fmt.Println(possibilities(".?"))
	fmt.Println(possibilities("?-?"))
}
