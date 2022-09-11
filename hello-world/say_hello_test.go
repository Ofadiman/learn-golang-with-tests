package main

import "testing"

func TestSayHello(t *testing.T) {
	got := SayHello()
	want := "Hello World"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
