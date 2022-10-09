package main

import "testing"

func assertEquals(t testing.TB, expected string, received string) {
	t.Helper()
	if expected != received {
		t.Errorf("expected %v to equal %v", expected, received)
	}
}

func assertNil(t testing.TB, expected any) {
	t.Helper()
	if expected != nil {
		t.Errorf("expected %v to equals nil", expected)
	}
}

func assertNotNil(t testing.TB, expected any) {
	t.Helper()
	if expected == nil {
		t.Errorf("expected %v not to equals nil", expected)
	}
}

func TestSayHello(t *testing.T) {
	const name = "Ofadiman"

	t.Run("should return greeting in english when english language is passed", func(t *testing.T) {
		got, err := SayHello(name, EnglishLanguage)
		want := EnglishHelloPrefix + name + HelloSuffix

		assertEquals(t, got, want)
		assertNil(t, err)
	})

	t.Run("should return greeting in spanish when spanish language is passed", func(t *testing.T) {
		got, err := SayHello(name, SpanishLanguage)
		want := SpanishHelloPrefix + name + HelloSuffix

		assertEquals(t, got, want)
		assertNil(t, err)
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		_, err := SayHello("", EnglishLanguage)

		assertNotNil(t, err)
	})

	t.Run("should return error when language is not supported", func(t *testing.T) {
		_, err := SayHello(name, "german")

		assertNotNil(t, err)
	})

}
