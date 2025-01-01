package interview

import (
	"fmt"
	"strings"
)

/*
	input  : aaabbaddd
	output : a3b2a1d3
*/

func CountCharecters(input string) string {

	if len(input) == 0 {
		return ""
	}

	var output strings.Builder
	currentChar := input[0]
	currentCharCount := 1

	for i := 1; i <= len(input)-1; i++ {

		if input[i] == currentChar {
			currentCharCount++
		} else {
			output.WriteString(fmt.Sprintf("%v%v", string(currentChar), currentCharCount))
			currentChar = input[i]
			currentCharCount = 1
		}
	}

	output.WriteString(fmt.Sprintf("%v%v", string(currentChar), currentCharCount))

	return output.String()
}
