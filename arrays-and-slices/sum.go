package main

import "errors"

func Sum(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("slice cannot be empty")
	}

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum, nil
}
