package main

import "errors"

func Repeat(character string, times int) (string, error) {
	if len(character) != 1 {
		return "", errors.New("string must be a single character")
	}

	repeated := ""
	for i := 0; i < times; i++ {
		repeated = repeated + character
	}

	return repeated, nil
}
