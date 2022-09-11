package main

import (
	"testing"
)

func assertArea(t testing.TB, shape Shape, expected float64) {
	t.Helper()
	result := shape.Area()
	if result != expected {
		t.Errorf("result %g not equals to %g", result, expected)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape       Shape
		expected    float64
		description string
	}{
		{shape: Rectangle{12, 6}, expected: 72.0, description: "should return rectangle area"},
		{shape: Circle{10}, expected: 314.1592653589793, description: "should return circle area"},
	}

	for _, test := range areaTests {
		t.Run(test.description, func(t *testing.T) {
			assertArea(t, test.shape, test.expected)
		})
	}
}
