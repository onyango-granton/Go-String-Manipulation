package go_reloaded

import (
	"regexp"
	//"strings"
)

// func isSpeechMark(s string) bool {
// 	speechMark := "'"
// 	return speechMark == s
// }

// func PairSpeechMark(s string) string {
// 	// words := strings.Split(s, " ")
// 	// isStartSpeechMark := false
// 	// for index, word := range words {
// 	// 	if isSpeechMark(word) {
// 	// 		isStartSpeechMark = !isStartSpeechMark
// 	// 		if isStartSpeechMark && index-1 > 0 && index+1 < len(words) {
// 	// 			words[index] = words[index] + words[index+1]
// 	// 			words = append(words[:index+1], words[index+2:]...)
// 	// 		} else if !isStartSpeechMark {
// 	// 			if index+1 < len(words) {
// 	// 				words[index-1] = words[index-1] + words[index]
// 	// 				words = append(words[:index], words[index+1:]...)
// 	// 			} else if index < len(words) {
// 	// 				words[index-1] = words[index-1] + words[index]
// 	// 				words = words[:index]
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }

// 	//return strings.Join(words, " ")

// }

func PairSpeechMark(s string) string {
	res := regexp.MustCompile(`'\s+(.*?)\s+'`)
	str2 := res.ReplaceAllString(s, "'$1'")
	return str2
}
