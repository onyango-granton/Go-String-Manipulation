package go_reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func ToUpper(s string) string {
	runeS := []rune(s)

	smallLetterStart := 97
	capitalLetterStart := 65

	newString := ""

	for i := 0; i < len(runeS); i++ {
		if runeS[i] >= 97 && runeS[i] <= 122 {
			diff := int(runeS[i]) - smallLetterStart
			newCapital := capitalLetterStart + diff

			newChar := string(rune(newCapital))

			newString = newString + newChar
		} else {
			newString = newString + string(rune(runeS[i]))
		}
	}

	return newString
}

func lastIndex(str string, s byte) int {
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == s {
			return i
		}
	}
	return -1
}

func Capitalize(s string) string {
	defaultValue := 1
	words := Split(s, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" {
			words[i-1] = ToUpper(string(words[i-1][0])) + words[i-1][1:]
		}
		if words[i] == "(cap," && i > 0 {
			if j := lastIndex(words[i+1], ')'); j != -1 {
				result := words[i+1][:j]
				num, err := strconv.Atoi(result)
				if err != nil {
					fmt.Println(err)
				}
				defaultValue = num
			}
			for x := 1; x <= defaultValue; x++ {
				words[i-x] = ToUpper(string(words[i-x][0])) + words[i-x][1:]
			}
		}
		if words[i] == "(cap)" && i > 0 {
			words[i] = ""
			words = append(words[:i], words[i+1:]...)
		}
		if words[i] == "(cap," && i > 0 {
			words[i] = ""
			words = append(words[:i], words[i+2:]...)

		}
	}

	// res := strings.Join(words, " ")
	// //re := regexp.MustCompile(`\s*\(\s*cap\s*\)\s*`)
	// re := regexp.MustCompile(`\s*\(\s*cap,\s*`+strconv.Itoa(defaultValue)+`\s*\)\s*`)
	// res = re.ReplaceAllString(res, "")
	// re2 := regexp.MustCompile(`\s*\(\s*cap\s*\)\s*`)
	// res = re2.ReplaceAllString(res, " ")
	// fmt.Println(defaultValue)

	return strings.Join(words, " ")

	// return strings.Join(words, " ")
}
