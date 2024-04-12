package go_reloaded

func ToLower(s string) string {
	runeS := []rune(s)

	smallLetterStart := 97
	capitalLetterStart := 65

	newString := ""

	for i := 0; i < len(runeS); i++ {
		if runeS[i] >= 65 && runeS[i] <= 90 {
			diff := int(runeS[i]) - capitalLetterStart
			newCapital := smallLetterStart + diff

			newChar := string(rune(newCapital))

			newString = newString + newChar
		} else {
			newString = newString + string(rune(runeS[i]))
		}
	}

	return newString
}
