package main

import "errors"

func SumTail(numberSlices ...[]int) (int, error) {
	if len(numberSlices) == 0 {
		return 0, errors.New("number slices cannot be empty")
	}

	tail := numberSlices[1:]
	if len(tail) == 0 {
		return 0, errors.New("tail must have at least 1 element")
	}

	sum := 0
	for _, numbers := range tail {
		for _, number := range numbers {
			sum += number
		}
	}

	return sum, nil
}
