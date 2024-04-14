package main

import "testing"

func TestManipulate(t *testing.T) {
	t.Run("Test with 1 string", func(t *testing.T) {
		input := []string{"a"}
		expected := []string{"a"}
		actual := Manipulate(input)
		if actual[0] != expected[0] {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})
	t.Run("Test with 2 strings", func(t *testing.T) {
		input := []string{"ab"}
		expected := []string{"ab", "ba"}
		actual := Manipulate(input)
		if actual[0] != expected[0] {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})
	t.Run("Test with 3 strings", func(t *testing.T) {
		input := []string{"abc"}
		expected := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
		actual := Manipulate(input)
		if actual[0] != expected[0] {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})
	t.Run("Test with 4 strings", func(t *testing.T) {
		input := []string{"aabb"}
		expected := []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}
		actual := Manipulate(input)
		if actual[0] != expected[0] {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
		if len(actual) != 6 {
			t.Errorf("Expected %d, got %d", len(expected), len(actual))
		}
	})
}
