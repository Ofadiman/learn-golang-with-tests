package main

import "testing"

func TestSayHello(t *testing.T) {
	const name = "Ofadiman"

	t.Run("should return greeting in english when english language is passed", func(t *testing.T) {
		got, err := SayHello(name, EnglishLanguage)
		want := EnglishHelloPrefix + name + HelloSuffix

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		if err != nil {
			t.Errorf("function SayHello returned an error")
		}
	})

	t.Run("should return greeting in spanish when spanish language is passed", func(t *testing.T) {
		got, err := SayHello(name, SpanishLanguage)
		want := SpanishHelloPrefix + name + HelloSuffix

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		if err != nil {
			t.Errorf("function SayHello returned an error")
		}
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		_, err := SayHello("", EnglishLanguage)

		if err == nil {
			t.Errorf("function SayHello did not return error")
		}
	})

	t.Run("should return error when language is not supported", func(t *testing.T) {
		_, err := SayHello(name, "german")

		if err == nil {
			t.Errorf("function SayHello did not return error")
		}
	})

}
