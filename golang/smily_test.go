package main

import "testing"

func TestCountSmilyFace(t *testing.T) {
	t.Run("should return 0 when input is empty", func(t *testing.T) {
		input := []string{}
		expected := 0
		actual := CountSmilyFace(input)
		if actual != expected {
			t.Errorf("got %d, want %d", actual, expected)
		}
	})
	t.Run("Return 2 smily faces", func(t *testing.T) {
		input := []string{":)", ";(", ";}", ":-D"}
		expected := 2
		actual := CountSmilyFace(input)
		if actual != expected {
			t.Errorf("got %d, want %d", actual, expected)
		}
	})
	t.Run("Return 3 smily faces", func(t *testing.T) {
		input := []string{";D", ":-(", ":-)", ";~)"}
		expected := 3
		actual := CountSmilyFace(input)
		if actual != expected {
			t.Errorf("got %d, want %d", actual, expected)
		}
	})
	t.Run("Return 1 smily faces", func(t *testing.T) {
		input := []string{";]", ":[", ";*", ":$", ";-D"}
		expected := 1
		actual := CountSmilyFace(input)
		if actual != expected {
			t.Errorf("got %d, want %d", actual, expected)
		}
	})
}
