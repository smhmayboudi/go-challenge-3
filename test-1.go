package main

import (
	"fmt"
	"math"
	"strings"
)

func generateCombinations(input string) {
	// Count the number of question marks in the input string
	numQuestionMarks := strings.Count(input, "?")

	// Generate all possible combinations
	for i := 0; i <= int(math.Pow(2, float64(numQuestionMarks)))-1; i++ {
		// Convert the current index to a binaryString string
		binaryString := fmt.Sprintf("%0[1]*[2]b", numQuestionMarks, i)
		binaryRune := []rune(binaryString)

		// Replace the question marks with the current binary string
		result := []rune(input)
		k := 0
		for j := 0; j < len(result); j++ {
			if result[j] != '?' {
				continue
			}
			result[j] = binaryRune[k]
			k++
		}

		// printout the result
		println(string(result))
	}
}

func Test1() {
	input := "1?101?"
	generateCombinations(input)
}
