package main

import (
	"reflect"
	"testing"
)

func assertDeepEqual(t testing.TB, expected any, received any) {
	t.Helper()
	if !reflect.DeepEqual(expected, received) {
		t.Errorf("got %v expected %v", expected, received)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("should return error when slices are empty", func(t *testing.T) {
		_, err := SumAll()

		assertNotNil(t, err)
	})

	t.Run("should sum slices one by one", func(t *testing.T) {
		result, err := SumAll([]int{1, 2, 3}, []int{4, 5, 6})

		expected := []int{6, 15}
		assertDeepEqual(t, expected, result)
		assertNil(t, err)
	})
}
