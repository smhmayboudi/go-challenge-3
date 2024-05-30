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

# Test 2

Write a client for redis which support to methods (Get, Set). It should support the time as well.

## First Try

```GO
package main

import "errors"

var ErrGet = errors.New("get error")
var ErrGetTime = errors.New("get error time")

// Serice is defined to support get and set.
type Service interface {
	Get(string, int) (string, error)
	Set(string, string, int) error
}

// Data is a struct to store information.
type Data struct {
	time  int
	value string
}

// ImplService is an implementation of Service interface.
type ImplService struct {
	// ds is a data source.
	ds map[string][]Data
}

// Get is a function.
func (is *ImplService) Get(key string, time int) (string, error) {
	val, ok := is.ds[key]
	if !ok {
		return "", ErrGet
	}
	mem := Data{}
	for i := 0; i < len(val); i++ {
		v := val[i]
		if v.time <= time {
			mem = v
		} else {
			break
		}
	}
	return mem.value, nil
}

// Set is a function.
func (is *ImplService) Set(key string, value string, time int) error {
	val, ok := is.ds[key]
	if !ok {
		is.ds[key] = make([]Data, 0)
	}
	if len(val) > 0 && val[len(val)-1].time > time {
		return ErrGetTime
	}
	newMem := append(val, Data{
		time:  time,
		value: value,
	})
	is.ds[key] = newMem
	return nil
}

func Test2() {
	test := ImplService{}
	test.ds = make(map[string][]Data)

	if err := test.Set("foo", "bar", 1); err != nil {
		println("errors: ", err)
		return
	}

	if err := test.Set("foo", "bar2", 10); err != nil {
		println("error: ", err)
		return
	}

	testOut1, err := test.Get("foo", 2)
	if err != nil {
		println("error: ", err)
		return
	}
	println(testOut1)

	testOut2, err := test.Get("foo", 1)
	if err != nil {
		println("error: ", err)
		return
	}
	println(testOut2)

	testOut3, err := test.Get("foo", 0)
	if err != nil {
		println("error: ", err)
		return
	}
	println(testOut3)

	testOut4, err := test.Get("foo", 11)
	if err != nil {
		println("error: ", err)
		return
	}
	println(testOut4)

	if err := test.Set("foo", "bar3", 10); err != nil {
		println("error: ", err)
		return
	}

	testOut5, err := test.Get("foo", 11)
	if err != nil {
		println("error: ", err)
		return
	}
	println(testOut5)

	if err := test.Set("foo", "bar4", 9); err != nil {
		println("error: ", err.Error())
		return
	}
}
```

## Better

```GO
package main

import (
	"errors"
	"sort"
)

var ErrNotFound = errors.New("key not found")
var ErrInvalidTimestamp = errors.New("invalid timestamp")

// Service is defined to support get and set.
type Service interface {
	Get(string, int) (string, error)
	Set(string, string, int) error
}

// Data is a struct to store information.
type Data struct {
	Timestamp int
	Value     string
}

// ImplService is an implementation of Service interface.
type ImplService struct {
	// ds is a data source.
	ds map[string][]Data
}

func NewImplService() *ImplService {
	return &ImplService{
		ds: make(map[string][]Data),
	}
}

// Get is a function.
func (is *ImplService) Get(key string, timestamp int) (string, error) {
	values, ok := is.ds[key]
	if !ok {
		return "", ErrNotFound
	}
	i := sort.Search(len(values), func(i int) bool {
		return values[i].Timestamp > timestamp
	})
	if i == 0 {
		return "", nil
	}
	return values[i-1].Value, nil
}

// Set is a function.
func (is *ImplService) Set(key string, value string, timestamp int) error {
	values, ok := is.ds[key]
	if !ok {
		values = make([]Data, 0)
	}
	if len(values) > 0 && values[len(values)-1].Timestamp > timestamp {
		return ErrInvalidTimestamp
	}
	is.ds[key] = append(values, Data{
		Timestamp: timestamp,
		Value:     value,
	})
	sort.Slice(is.ds[key], func(i, j int) bool {
		return is.ds[key][i].Timestamp < is.ds[key][j].Timestamp
	})
	return nil
}

func Test2() {
	test := NewImplService()

	if err := test.Set("foo", "bar", 1); err != nil {
		println("errors: ", err.Error())
		return
	}

	if err := test.Set("foo", "bar2", 10); err != nil {
		println("error: ", err.Error())
		return
	}

	testOut1, err := test.Get("foo", 2)
	if err != nil {
		println("error: ", err.Error())
		return
	}
	println(testOut1)

	testOut2, err := test.Get("foo", 1)
	if err != nil {
		println("error: ", err.Error())
		return
	}
	println(testOut2)

	testOut3, err := test.Get("foo", 0)
	if err != nil {
		println("error: ", err.Error())
		return
	}
	println(testOut3)

	testOut4, err := test.Get("foo", 11)
	if err != nil {
		println("error: ", err.Error())
		return
	}
	println(testOut4)

	if err := test.Set("foo", "bar3", 10); err != nil {
		println("error: ", err.Error())
		return
	}

	testOut5, err := test.Get("foo", 11)
	if err != nil {
		println("error: ", err.Error())
		return
	}
	println(testOut5)

	if err := test.Set("foo", "bar4", 9); err != nil {
		println("error: ", err.Error())
		return
	}
}
```

The main changes in this version are:

- Error Handling: The custom error types have been replaced with more descriptive errors, ErrNotFound and ErrInvalidTimestamp.
- Efficiency of Get() Function: The Get() function now uses a binary search to find the most recent value before the requested timestamp, improving the time complexity from O(n) to O(log n), where n is the number of values for the given key.
- Concurrency: The current implementation is still not thread-safe, but you can easily add a mutex to protect the shared data map if needed.
- Test Coverage: The provided main() function serves as a basic test case, but you should consider adding more comprehensive tests to cover all edge cases.
