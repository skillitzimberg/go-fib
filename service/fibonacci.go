package main

import (
	"errors"
	"fmt"
	"math"
)

var err error

func fib(num int) ([]uint64, error) {
	result := []uint64{0}
	var a, b uint64 = 0, 1
	for len(result) < num {
		if a+b > math.MaxInt64 {
			err = errors.New("server encountered an integer overflow")
			if err != nil {
				fmt.Print(err)
			}
			return nil, err
		}
		a, b = b, a+b
		result = append(result, a)

	}
	return result, err
}
