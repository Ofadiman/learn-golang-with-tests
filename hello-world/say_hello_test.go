package main

import "testing"

func TestSayHello(t *testing.T) {
	t.Run("should return hello string", func(t *testing.T) {
		const name = "Ofadiman"
		got, err := SayHello(name)
		want := EnglishHelloPrefix + name

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		if err != nil {
			t.Errorf("function SayHello returned an error")
		}
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		_, err := SayHello("")

		if err == nil {
			t.Errorf("function SayHello did not return error")
		}
	})

}
