package main

import "testing"

func TestFindOddNumber(t *testing.T) {
	t.Run("Test with 1 element", func(t *testing.T) {
		input := []int{7}
		expected := 7
		actual := FindOddNumber(input)
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
	t.Run("Test with 1 element with zero", func(t *testing.T) {
		input := []int{0}
		expected := 0
		actual := FindOddNumber(input)
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
	t.Run("Test with 3 elements", func(t *testing.T) {
		input := []int{1, 1, 2}
		expected := 2
		actual := FindOddNumber(input)
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
	t.Run("Test with 4 elements", func(t *testing.T) {
		input := []int{0, 1, 0, 1, 0}
		expected := 0
		actual := FindOddNumber(input)
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
	t.Run("Test with 11 elements", func(t *testing.T) {
		input := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
		expected := 4
		actual := FindOddNumber(input)
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}
