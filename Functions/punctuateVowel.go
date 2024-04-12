package go_reloaded

import "strings"

func isVowelWord(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "h"}
	
	for _, vowel := range vowels{
		if vowel == string(s[0]) && len(s) > 1{
			return true
		}		
	}
	return false
}

func PunctuateVowel(s string) string {
	words := strings.Split(s, " ")
	for index, word := range words{
		if isVowelWord(word) && index > 0 && index + 1 < len(words) {
			if words[index - 1] == "a" || words[index - 1] == "A"{
				words[index - 1] = words[index - 1] + "n"
			}
		}  
	}
	return strings.Join(words, " ")
}