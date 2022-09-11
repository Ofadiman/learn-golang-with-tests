package main

import "errors"

type Dictionary map[string]string

var ElementNotFoundError = errors.New("ElementNotFoundError")
var ElementAlreadyExistsError = errors.New("ElementAlreadyExistsError")

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]

	if ok == false {
		return "", ElementNotFoundError
	}

	return result, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, ok := d[word]
	if ok == true {
		return ElementAlreadyExistsError
	}

	d[word] = definition
	return nil
}
