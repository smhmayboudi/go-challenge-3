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
