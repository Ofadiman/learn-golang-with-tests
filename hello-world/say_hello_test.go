package main

import "testing"

func TestSayHello(t *testing.T) {
	const name = "Ofadiman"
	got := SayHello(name)
	want := EnglishHelloPrefix + name

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
