package go_reloaded

import "strings"

func isGroupPunctations(s string) bool {
	marks := []string{".", ",", "!", "?", ":", ";"}
	count := 0
	alphaMarks := []bool{}

	for _, alpha := range s {
		for _, mark := range marks {
			if mark == string(alpha) {
				alphaMarks = append(alphaMarks, true)
			} else {
				alphaMarks = append(alphaMarks, false)
			}
		}
	}

	for _, alphaMark := range alphaMarks {
		if alphaMark {
			count++
		}
	}

	return count > 1
}

func GroupPunctations(s string) string {
	words := Split(s, " ")
	for index, word := range words {
		if isGroupPunctations(word) {
			if index+1 < len(words) {
				words[index-1] = words[index-1] + words[index]
				words = append(words[:index], words[index+1:]...)
			} else if index < len(words) {
				words[index-1] = words[index-1] + words[index]
				words = words[:index]
			}
		}
	}
	return strings.Join(words, " ")
}
