package main

func Manipulate(text []string) []string {
	var permutations []string
	seen := make(map[string]bool)

	var generate func(chars []rune, index int)
	generate = func(chars []rune, index int) {
		if index == len(chars) {
			perm := string(chars)
			if !seen[perm] {
				permutations = append(permutations, perm)
				seen[perm] = true
			}
			return
		}

		for i := index; i < len(chars); i++ {
			chars[index], chars[i] = chars[i], chars[index]
			generate(chars, index+1)
			chars[index], chars[i] = chars[i], chars[index]
		}
	}

	for _, str := range text {
		chars := []rune(str)
		generate(chars, 0)
	}

	return permutations
}
