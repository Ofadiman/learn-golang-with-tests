package main

import "testing"

func assertEquals(t testing.TB, expected any, received any) {
	t.Helper()
	if expected != received {
		t.Errorf("expected %v to equal %v", expected, received)
	}
}

func TestSearch(t *testing.T) {
	const key = "800757b7-aca1-4c0a-8912-7f57c17937d3"
	const value = "this is just a test"
	dictionary := Dictionary{key: value}

	t.Run("existing key", func(t *testing.T) {
		result, _ := dictionary.Search(key)

		assertEquals(t, result, value)
	})

	t.Run("non-existing key", func(t *testing.T) {
		_, err := dictionary.Search("whatever")

		assertEquals(t, err, ElementNotFoundError)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding non-existing word", func(t *testing.T) {
		const word = "test"
		const definition = "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Add(word, definition)
		assertEquals(t, err, nil)

		result, err := dictionary.Search(word)
		assertEquals(t, err, nil)
		assertEquals(t, result, definition)
	})

	t.Run("adding existing word", func(t *testing.T) {
		const word = "test"
		const definition = "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)
		assertEquals(t, err, ElementAlreadyExistsError)
	})

}
