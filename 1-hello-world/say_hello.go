package main

import (
	"errors"
)

const EnglishHelloPrefix = "Hello, "
const EnglishLanguage = "english"

const SpanishLanguage = "spanish"
const SpanishHelloPrefix = "Hola, "

const HelloSuffix = "!"

var languageMap = map[string]string{
	EnglishLanguage: EnglishHelloPrefix,
	SpanishLanguage: SpanishHelloPrefix,
}

func SayHello(name string, language string) (string, error) {
	prefix, ok := languageMap[language]
	if ok == false {
		return "", errors.New("language " + language + " is not supported")
	}

	if name == "" {
		return "", errors.New("string cannot be empty")
	}

	return prefix + name + HelloSuffix, nil
}
