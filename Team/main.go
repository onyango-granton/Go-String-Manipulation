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

func isSpeechMark(s string) bool {
	speechMark := "'"
	if speechMark == s{
		return true
	}
	return false
}

// func pairSpeechMark(s string) string {
// 	words := strings.Split(s, " ")
// 	for i, word := range(words){
// 		if isSpeechMark(word) && words[i - 2] != "'" && i-2 >= 0 && i + 2 < len(words){
// 			words[i] = words[i] + words[i+1] + words[i+2]
// 			words = append(words[:i + 1], words[i + 3:]...)
			
// 		}
// 	}
// 	return strings.Join(words, " ")
// }

func isVowelWord(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "h"}
	
	for _, vowel := range vowels{
		if vowel == string(s[0]) && len(s) > 1{
			return true
		}		
	}
	return false
}

func punctuateVowel(s string) string {
	words := strings.Split(s, " ")
	for index, word := range words{
		if isVowelWord(word) && index > 0 && index + 1 < len(words) && words[index - 1] == "a"{
			words[index - 1] = "an"
			//words = append(words[:index - 1], words[index :]...)
		}  
	}
	return strings.Join(words, " ")
}

func isStartingWithPunc(s string) bool {
	marks := []string{".", ",", "!", "?", ":"}
	for _, alpha := range s{
		for _, mark := range marks{ 
			if mark == string(alpha) && len(s) > 1{
				return true
			}
		}
	}
	return false
}

func formatSentenceWordStrtPunc(s string) string {
	words := strings.Split(s, " ")
	for index, word := range words{
		if isStartingWithPunc(word) && index > 0 {
			words[index - 1] = words[index - 1] + string(word[0])
			words[index] = words[index][1:]
		}
	}
	return strings.Join(words, " ")
}

func pairSpeechMark(s string) string {
	words := strings.Split(s, " ")
	isStartSpeechMark := false
	for index, word := range words{
		if isSpeechMark(word){
			isStartSpeechMark = !isStartSpeechMark
			if isStartSpeechMark && index - 1 > 0 && index + 1 < len(words){
				words[index] = words[index] + words[index + 1]
				words = append(words[:index + 1], words[index + 2:]...) 
			} else if !isStartSpeechMark && index - 1 > 0 {
				// words[index-1] = words[index - 1] + words[index]
				// words = append(words[:index], words[index + 1:]...)
				if index + 1 < len(words){
					words[index-1] = words[index - 1] + words[index]
					words = append(words[:index], words[index + 1:]...)
				} else if index < len(words){
					words[index-1] = words[index - 1] + words[index]
					words = words[:index]
					fmt.Println("Im Here")
				}

			}
		}
	}

	return strings.Join(words, " ")
}


func main() {
	//s1 := "Hello . world"
	//s1 := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	s1 := "a ,amazing ,day '"
	s2 := "Today I said ' veni vidi vici '"
	s3 := "LIving like we are ' renegades '"
	//s2 := strings.Split(s1, " ")
	// for _, word := range s2 {
	// 	fmt.Println(isSpecialword(word))
	// }
	fmt.Println(formatSentenceWordStrtPunc(s1))
	fmt.Println(pairSpeechMark(s2))
	fmt.Println(pairSpeechMark(s3))
}
