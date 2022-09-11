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

func TestRepeat(t *testing.T) {
	t.Run("should return error when an empty string is passed", func(t *testing.T) {
		result, err := Repeat("", 5)

		assertEquals(t, result, "")
		assertNotNil(t, err)
	})

	t.Run("should return error when string longer than 1 character is passed", func(t *testing.T) {
		result, err := Repeat("ab", 5)

		assertEquals(t, result, "")
		assertNotNil(t, err)
	})

	t.Run("should repeat character N number of times", func(t *testing.T) {
		result, err := Repeat("a", 5)

		assertEquals(t, result, "aaaaa")
		assertNil(t, err)
	})
}
