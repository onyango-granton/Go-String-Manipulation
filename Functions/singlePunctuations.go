package go_reloaded

import (
	"strings"
)

func isStartingWithPunc(s string) bool {
	marks := []string{".", ",", "!", "?", ":"}
	for index, alpha := range s{
		for _, mark := range marks{ 
			if mark == string(alpha) && index == 0{
				return true
			}
		}
	}
	return false
}

 func SinglePunctuations(s string) string {
	marks := []string{".", ",", "!", "?", ":" ,";"}
	words := strings.Split(s, " ")
	

	for index, word := range words{
		for _, mark := range marks{
			if word == mark {
				words[index - 1] = words[index - 1] + words[index]
				words = append(words[:index], words[index+1:]...)
			}
		}
		// if isStartingWithPunc(word) && index > 0 && index < len(words){
		// 	words[index - 1] = words[index - 1] + string(word[0])
		// 	words[index] = words[index][1:]
		// 	words = append(words[:index],words[index:]...)
		// }
	}


	
	return strings.Join(words, " ")

	
 }

 func WordStartPunc(s string) string {
	words := strings.Split(s, " ")
	for index, word := range words{
		if isStartingWithPunc(word) && index > 0 && index < len(words){
			words[index - 1] = words[index - 1] + string(word[0])
			words[index] = words[index][1:]
			words = append(words[:index],words[index:]...)
		}
	}
	return strings.Join(words, " ")
 }