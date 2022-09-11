package main

import "testing"

func TestSumTail(t *testing.T) {
	t.Run("should return error when passed list is empty", func(t *testing.T) {
		_, err := SumTail()

		assertNotNil(t, err)
	})

	t.Run("should return error when passed list has only 1 element", func(t *testing.T) {
		_, err := SumTail([]int{1, 2, 3})

		assertNotNil(t, err)
	})

	t.Run("should return tail sum", func(t *testing.T) {
		result, err := SumTail([]int{1, 2, 3}, []int{4, 5, 6})

		assertEquals(t, 15, result)
		assertNil(t, err)
	})
}
