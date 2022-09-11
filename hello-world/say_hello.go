package main

import "errors"

const EnglishHelloPrefix = "Hello, "

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("string cannot be empty")
	}
	return EnglishHelloPrefix + name, nil
}
