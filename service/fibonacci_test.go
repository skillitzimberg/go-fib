package main

import (
	"reflect"
	"testing"
)

type test struct {
	input  int
	output []uint64
}

var tests = []test{
	{1, []uint64{0}},
	{2, []uint64{0, 1}},
	{3, []uint64{0, 1, 1}},
	{4, []uint64{0, 1, 1, 2}},
}

func TestFibonnaci(t *testing.T) {
	for _, test := range tests {
		if reflect.TypeOf(test.input) != reflect.TypeOf(1) {
			t.Errorf("Input must be an integer.")
		}

		if test.input < 0 {
			t.Errorf("Input must be greater than 0.")
		}
		expected := test.output
		received, err := fib(test.input)
		check(err, "test : fib(num)")
		for i, num := range expected {
			if received[i] != num {
				t.Errorf("Expected %v, got %v", received[i], num)
			}

		}
	}
}
