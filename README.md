# go-challenge-3

# Test 1

Write a function that takes a string as input and generates all possible combinations of the string by replacing the question marks (?) with 0 or 1.

## First Try

```GO
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Name(in string) {
	numQ := strings.Count(in, "?")
	for i := 0; i <= int(math.Pow(2, float64(numQ)))-1; i++ {
		binary := strconv.FormatInt(int64(i), 2)
		str := fmt.Sprintf("%0"+fmt.Sprint(numQ)+"b", binary)
		str = strings.ReplaceAll(str, "%!b(string=", "")
		s := []rune(str)
		a := []rune(in)
		l := len(a)
		k := 0
		for j := 0; j < l; j++ {
			if a[j] == '?' {
				a[j] = s[k]
				k++
			}
		}
		println(string(a))
	}

}

func main() {
	Name("1?101?")
}
```

## Better

```GO
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

func main() {
	input := "1?101?"
	generateCombinations(input)
}
```

Here's how the generateCombinations() function works:

- The number of question marks in the input string is counted using the strings.Count() function.
- A loop is used to iterate through all possible combinations, where the number of iterations is 2^n, where n is the number of question marks.
- For each iteration, the current index is converted to a binary string using fmt.Sprintf().
- A new slice of runes is created to hold the resulting combination. The function then iterates through the input string, replacing the question marks with the corresponding binary digits.
- The resulting combination is then printed using fmt.Println().

The time complexity of this solution is O(2^n), where n is the number of question marks in the input string, as it generates all possible combinations.
