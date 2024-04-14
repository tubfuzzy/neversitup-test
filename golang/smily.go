package main

import "regexp"

func CountSmilyFace(text []string) int {
	count := 0
	pattern := `^[:;][-~]?[)D]$`
	for _, str := range text {
		if regexp.MustCompile(pattern).MatchString(str) {
			count++
		}
	}

	return count
}
