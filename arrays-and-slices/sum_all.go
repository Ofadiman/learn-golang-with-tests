package main

import "errors"

func SumAll(numberSlices ...[]int) ([]int, error) {
	numberSlicesCount := len(numberSlices)

	if numberSlicesCount == 0 {
		return []int{}, errors.New("slices cannot be empty")
	}

	var sums []int
	for _, numberSlice := range numberSlices {
		currentSliceSum := 0
		for _, number := range numberSlice {
			currentSliceSum += number
		}
		sums = append(sums, currentSliceSum)
	}

	return sums, nil
}
