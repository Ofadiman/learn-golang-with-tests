package main

import "testing"

func assertEquals(t testing.TB, expected any, received any) {
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

func TestSum(t *testing.T) {
	t.Run("should return error when slice is empty", func(t *testing.T) {
		numbers := []int{}
		_, err := Sum(numbers)

		assertNotNil(t, err)
	})

	t.Run("should sum numbers from slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result, err := Sum(numbers)

		assertNil(t, err)
		assertEquals(t, result, 6)
	})
}
