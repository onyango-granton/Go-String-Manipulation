package main

import (
	"fmt"
	"strings"
)

func isSpecialword(s string) bool {
	count := 0
	boolChar := []bool{}
	characters := []string{".", ",", "!", "?", ":"}
	for _, word := range s {
		for _, c := range characters {
			if c == string(word) {
				boolChar = append(boolChar, true)
			} else {
				boolChar = append(boolChar, false)
			}
		}
	}

	for _, bc := range boolChar {
		if bc {
			count++
		}
	}

	if count > 1 {
		return true
	}

	return false
}

func formatSentence(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		if isSpecialword(word) && i > 0 {
			//words[i - 1] = ""
			words[i-1] = words[i-1] + words[i]
			words = append(words[:i], words[i+1:]...)
		}
	}

	s2 := strings.Join(words, " ")

	return s2
}

func isMark(s string) bool {
	marks := []string{".", ",", "!", "?", ":"}
	for _, mark := range marks {
		if mark == s {
			return true
		}
	}
	return false
}

func formatPunctuation(word string) string {
	words := strings.Split(word, " ")
	for i, word := range words{
		if isMark(word) && i > 0 {
			words[i - 1] = words[i - 1] + words[i]
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}

//func isSpeec

func main() {
	//s1 := "Hello . world"
	s1 := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	s2 := strings.Split(s1, " ")
	for _, word := range s2 {
		fmt.Println(isSpecialword(word))
	}
	fmt.Println(formatPunctuation(s1))
}
