package go_reloaded

func OutputFunc(s1 string) string {
	s1 = BinToDec(s1)
	s1 = HexToDec(s1)
	s1 = GroupPunctations(s1)
	s1 = SinglePunctuations(s1)
	s1 = Capitalize(s1)
	s1 = UpperCase(s1)
	s1 = LowerCase(s1)
	//s1 = GroupPunctations(s1)
	//s1 = SinglePunctuations(s1)
	s1 = WordStartPunc(s1) //
	s1 = PairSpeechMark(s1)
	s1 = PunctuateVowel(s1)

	return s1
}