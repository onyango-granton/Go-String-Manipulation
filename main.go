// package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strconv"
// 	"strings"
// )

// func Split(s string, sep string) []string {
// 	var result []string
// 	start := 0
// 	for i := 0; i < len(s); i++ {
// 		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
// 			result = append(result, s[start:i])
// 			start = i + len(sep)
// 			i = start - 1
// 		}
// 	}
// 	result = append(result, s[start:])
// 	return result
// }

// func Capitalize(s string) string {
// 	s = ToLower(s)
// 	for i, v := range s {
// 		if i == 0 {
// 			s = ToUpper(string(v)) + s[i+1:]
// 		} else {
// 			if IsAlpha(string(v)) && !IsAlpha(string(s[i-1])) {
// 				if i != len(s)-1 {
// 					s = s[:i] + ToUpper(string(v)) + s[i+1:]
// 				} else {
// 					s = s[:i] + ToUpper(string(v))
// 				}
// 			}
// 		}
// 	}
// 	return s
// }

// func IsAlpha(s string) bool {
// 	runeS := []rune(s)

// 	for i := 0; i < len(runeS); i++ {
// 		if runeS[i] >= 0 && runeS[i] < 48 || runeS[i] > 57 && runeS[i] < 65 || runeS[i] > 90 && runeS[i] < 97 || runeS[i] > 122 && runeS[i] <= 127 {
// 			return false
// 		}
// 	}

// 	return true
// }

// func ToLower(s string) string {
// 	runeS := []rune(s)

// 	smallLetterStart := 97
// 	capitalLetterStart := 65

// 	newString := ""

// 	for i := 0; i < len(runeS); i++ {
// 		if runeS[i] >= 65 && runeS[i] <= 90 {
// 			diff := int(runeS[i]) - capitalLetterStart
// 			newCapital := smallLetterStart + diff

// 			newChar := string(rune(newCapital))

// 			newString = newString + newChar
// 		} else {
// 			newString = newString + string(rune(runeS[i]))
// 		}
// 	}

// 	return newString
// }

// func ToUpper(s string) string {
// 	runeS := []rune(s)

// 	smallLetterStart := 97
// 	capitalLetterStart := 65

// 	newString := ""

// 	for i := 0; i < len(runeS); i++ {
// 		if runeS[i] >= 97 && runeS[i] <= 122 {
// 			diff := int(runeS[i]) - smallLetterStart
// 			newCapital := capitalLetterStart + diff

// 			newChar := string(rune(newCapital))

// 			newString = newString + newChar
// 		} else {
// 			newString = newString + string(rune(runeS[i]))
// 		}
// 	}

// 	return newString
// }

// // func hTD(word string) int {
// // 	keyValue := map[string]int{
// // 		"0" : 0,
// // 		"1" : 1,
// // 		"2" : 2,
// // 		"3" : 3,
// // 		"4" : 4,
// // 		"5" : 5,
// // 		"6" : 6,
// // 		"7" : 7,
// // 		"8" : 8,
// // 		"9" : 9,
// // 		"A" : 10,
// // 		"B" : 11,
// // 		"C" : 12,
// // 		"D" : 13,
// // 		"E" : 14,
// // 		"F" : 15,

// // 	}

// // 	//sampleString := "Hello World"

// // 	//sampleString := "1E"

// // 	var total int = 0

// // 	for k,v := range word {
// // 		index := len(word) - 1 - k
// // 		//fmt.Println(index,string(v))
// // 		value := float64(keyValue[string(v)]) * math.Pow(16 , float64(index))
// // 		total += int(value)
// // 		//fmt.Println(float64(keyValue[string(v)]) * math.Pow(16 , float64(index)))
// // 	}

// // 	return total
// // }

// func htD(s string) int64 {
// 	dec, err := strconv.ParseInt(s, 16, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return dec
// }

// func editor(s string) string {
// 	res := ""
// 	var s1 string = "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."

// 	s2 := ToLower(s1)

// 	// fmt.Println(s2)

// 	splitS1 := Split(s2, " ")

// 	fmt.Println(splitS1)

// 	for i := 0; i < len(splitS1); i++ {
// 		if strings.HasSuffix(splitS1[i], "(cap)") {
// 			res += Capitalize(splitS1[i-1])
// 		}
// 		res += splitS1[i]
// 	}

// 	return res
// }

// func capitalizeBeforeCap(sentence string) string {
//     //words := strings.Fields(sentence)
// 	words := Split(sentence, " ")// Split the sentence into words
//     for i := 0; i < len(words); i++ {
//         if words[i] == "(cap)" && i > 0 {
//             // If "(cap)" is found and it's not the first word
//             words[i-1] = Capitalize(words[i-1]) // Convert the word before "(cap)" to uppercase
//         }
//     }
//     return strings.Join(words, " ") // Join the words back into a sentence
// }

// func uppercaseBeforeUp(s string, optionalParam ...int) string {
// 	defaultValue := 1
// 	if (len(optionalParam) > 0){
// 		defaultValue = optionalParam[0]
// 	}
// 	words := Split(s, " ")
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(up)" && i > 0 {
// 			for j := 1; j <= defaultValue; j++ {
// 				words[i-j] = ToUpper(words[i-j])
// 			}

// 			//words[i-1] = ToUpper(words[i-1])
// 			//words[i] = ""
// 		}
// 	}
// 	return strings.Join(words, " ")
// }

//  func numberAfterUp(s string) string {
// 	defaultValue := 1
// 	words := Split(s, " ")
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(up)"{
// 			words[i-1] = ToUpper(words[i-1])
// 		}
// 		if words[i] == "(up," && i > 0 {
// 			if j := lastIndex(words[i + 1], ')'); j != -1{
// 				result := words[i+1][:j]
// 				//fmt.Println(result)
// 				num, err := strconv.Atoi(result)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 				defaultValue = num
// 			}
// 			//fmt.Println(defaultValue)
// 			for x := 1; x <= defaultValue; x++ {
// 				words[i-x] = ToUpper(words[i-x])
// 			}
// 		}
// 		// if words[i] == "(up)" && i > 0 {
// 		// 	words[i] = ""
// 		// }
// 		// if words[i] == "(up," && i > 0 {
// 		// 	words[i] = ""
// 		// 	words[i + 1] = ""

// 		// }
// 	}
// 	res := strings.Join(words, " ")
// 	//re := regexp.MustCompile(`\s*\(\s*cap\s*\)\s*`)
// 	re := regexp.MustCompile(`\s*\(\s*up,\s*`+strconv.Itoa(defaultValue)+`\s*\)\s*`)
// 	res = (re.ReplaceAllString(res, ""))
// 	re2 := regexp.MustCompile(`\s*\(\s*up\s*\)\s*`)
// 	res = (re2.ReplaceAllString(res, " "))
// 	fmt.Println(defaultValue)

// 	return res
//  }

//  func numberAfterLow(s string) string {
// 	defaultValue := 1
// 	words := Split(s, " ")
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(low)"{
// 			words[i-1] = ToUpper(words[i-1])
// 		}
// 		if words[i] == "(low," && i > 0 {
// 			if j := lastIndex(words[i + 1], ')'); j != -1{
// 				result := words[i+1][:j]
// 				//fmt.Println(result)
// 				num, err := strconv.Atoi(result)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 				defaultValue = num
// 			}
// 			//fmt.Println(defaultValue)
// 			for x := 1; x <= defaultValue; x++ {
// 				words[i-x] = ToUpper(words[i-x])
// 			}
// 		}
// 		if words[i] == "(low)" && i > 0 {
// 			words[i] = ""
// 		}
// 		if words[i] == "(low," && i > 0 {
// 			words[i] = ""
// 			words[i + 1] = ""

// 		}
// 	}

// 	res := strings.Join(words, " ")
// 	//re := regexp.MustCompile(`\s*\(\s*cap\s*\)\s*`)
// 	re := regexp.MustCompile(`\s*\(\s*low,\s*`+strconv.Itoa(defaultValue)+`\s*\)\s*`)
// 	res = (re.ReplaceAllString(res, ""))
// 	re2 := regexp.MustCompile(`\s*\(\s*low\s*\)\s*`)
// 	res = (re2.ReplaceAllString(res, " "))
// 	fmt.Println(defaultValue)

// 	return res

//  }

//  func numberAfterCap(s string) string {
// 	defaultValue := 1
// 	words := Split(s, " ")
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(cap)"{
// 			words[i-1] = ToUpper(words[i-1])
// 		}
// 		if words[i] == "(cap," && i > 0 {
// 			if j := lastIndex(words[i + 1], ')'); j != -1{
// 				result := words[i+1][:j]
// 				//fmt.Println(result)
// 				num, err := strconv.Atoi(result)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 				defaultValue = num
// 			}
// 			//fmt.Println(defaultValue)
// 			for x := 1; x <= defaultValue; x++ {
// 				words[i-x] = ToUpper(words[i-x])
// 			}
// 		}
// 		if words[i] == "(cap)" && i > 0 {
// 			words[i] = ""
// 		}
// 		if words[i] == "(cap," && i > 0 {
// 			words[i] = ""
// 			words[i + 1] = ""

// 		}
// 	}

// 	res := strings.Join(words, " ")
// 	//re := regexp.MustCompile(`\s*\(\s*cap\s*\)\s*`)
// 	re := regexp.MustCompile(`\s*\(\s*cap,\s*`+strconv.Itoa(defaultValue)+`\s*\)\s*`)
// 	res = re.ReplaceAllString(res, "")
// 	re2 := regexp.MustCompile(`\s*\(\s*cap\s*\)\s*`)
// 	res = re2.ReplaceAllString(res, " ")
// 	fmt.Println(defaultValue)

// 	return res

// 	//return strings.Join(words, " ")
//  }

//  func punctuations(s string) {
// 	marks := []string{".", ",", "!", "?", ":" ,";"}

// 	for _,j := range marks{
// 		k := ""

// 		if j == "?" {
// 			k = "\\?"
// 		} else if j == "." {
// 			k = "\\."
// 		} else {
// 			k = j
// 		}

// 		re := regexp.MustCompile(`\s*`+k+`\s*`)
// 		s = re.ReplaceAllString(s, j + " ")
// 	}

// 	fmt.Println(s)
//  }

// //  func followingPunc(s string) {
// // 	words := Split(s, " ")
// // 	marks := []string{".", ",", "!", "?", ":" ,";"}

// // 	markPresence := make(map[string]bool)

// // 	for _, mark := range marks{
// // 		markPresence[mark] = false
// // 	}

// // 	for _, word := range words{
// // 		count := 0
// // 		alphaMark := []bool{}
// // 		for _, alpha := range word{

// // 			for _, mark := range marks{
// // 				if mark == string(alpha){
// // 					alphaMark = append(alphaMark, true)
// // 				} else {
// // 					alphaMark = append(alphaMark, false)
// // 				}
// // 			}

// // 		}
// // 		for _, aM := range alphaMark{
// // 		 	if aM {
// // 		 		count++
// // 		 	}
// // 		 }
// // 		fmt.Println(count)
// // 	}

// //  }
//  func isFollowingPunc(s string) bool {
// 	marks := []string{".", ",", "!", "?", ":" ,";"}
// 	count := 0
// 	//markPresence := make(map[string]bool)
// 	alphaMarks := []bool{}

// 	for _, alpha := range s{
// 		for _, mark := range marks{
// 			if mark == string(alpha){
// 				alphaMarks = append(alphaMarks, true)
// 			} else {
// 				alphaMarks = append(alphaMarks, false)
// 			}
// 		}
// 	}

// 	for _, alphaMark := range alphaMarks{
// 		if alphaMark{
// 			count++
// 		}
// 	}

// 	if count > 1{
// 		return true
// 	}

// 	return false

//  }

//  func followingPunc(s string) string {
// 	s2 := ""
// 	s3 := ""
// 	s1 := Split(s, " ")
// 	sSlice := []string{}
// 	for _, word := range s1{
// 		if isFollowingPunc(word){
// 			s2 = strings.Join(s1, " ")
// 			re := regexp.MustCompile(`\s*`+word+`\s*`)
// 			s3 = re.ReplaceAllString(s2, word + " ")
// 		}
// 		sSlice = append(sSlice, word)
// 	}
// 	return s3
//  }

//  func lastIndex(str string, s byte) int {
//     for i := len(str) - 1; i >= 0; i-- {
//         if str[i] == s {
//             return i
//         }
//     }
//     return -1
// }

// func lowercaseBeforeLow(s string) string {
// 	words := Split(s, " ")
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(low)" && i > 0 {
// 			words[i-1] = ToLower(words[i-1])
// 			//words[i] = ""
// 		}
// 	}
// 	return strings.Join(words, " ")
// }

// func decimalBeforeHex(s string) string {
// 	words := Split(s, " ")
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(hex)" && i > 0 {
// 			dec, err := strconv.ParseInt(words[i - 1], 16, 64)
// 			if err != nil{
// 				fmt.Println(err)
// 			}
// 			words[i-1] = strconv.Itoa(int(dec))
// 			words[i-1] = ToLower(words[i-1])
// 			//words[i] = ""
// 		}
// 	}
// 	return strings.Join(words, " ")
// }

// func cleanUpWhiteSpace(s string) {
// 	words := Split(s, " ")
// 	for i:= 0; i < len(words); i++{

// 		if words[i] == "" {
// 			words[i] = "hello2"
// 		}
// 		if words[i] == " " && words[i+1] == " " && words[i+2] != " " {
// 			words[i] = ""
// 			words[i + 1] = ""
// 		}
// 		fmt.Println(words[i])
// 	}
// 	//fmt.Println(strings.Join(words, " "))
// }

// func decimalBeforeBin(s string) string {
// 	words := Split(s, " ")
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(bin)" && i > 0 {
// 			dec, err := strconv.ParseInt(words[i - 1], 2, 64)
// 			if err != nil{
// 				fmt.Println(err)
// 			}
// 			words[i-1] = strconv.Itoa(int(dec))
// 			words[i-1] = ToLower(words[i-1])
// 			//words[i] = ""
// 		}
// 	}
// 	return strings.Join(words, " ")
// }

// func removeBlockWhiteSpace(s string) string{
// 	words := Split(s, " ")
// 	for i := 0; i < len(words); i++{
// 		if words[i] == "(cap)" && i > 0 {
// 			words[i] = ""

// 		}
// 	}
// 	return strings.Join(words, " ")
// }

// func main() {
// 	//var s1 string = "it  (cap) was the best of (cap, 3) times  ., ,, !, ?, : and ; , it was the worst of times (up, 3) , it was the age of (up) wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was THE EPOCH OF INCREDULITY (low, 3) , it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."

// 	var s2 string = "Hello .... world"
// 	//var s2 string = "1E (hex) files were added"

// 	//var s3 string = "It has been 10 (bin) years"

// 	//fmt.Println(capitalizeBeforeCap(s1))

// 	//fmt.Println(uppercaseBeforeUp(s1, 2))

// 	//fmt.Println(lowercaseBeforeLow(s1))

// 	//fmt.Println(decimalBeforeHex(s2))

// 	//fmt.Println(decimalBeforeBin(s3))

// 	//fmt.Println(numberAfterUp(s1))

// 	//punctuations(s1)

// 	// s3 := Split(s2, " ")

// 	// for _, word := range s3{
// 	// 	fmt.Println(isFollowingPunc(word))
// 	// }

// 	fmt.Println(followingPunc(s2))

// 	//fmt.Println(numberAfterLow(s1))

// 	//fmt.Println(numberAfterCap(s1))
// 	//fmt.Println(removeBlockWhiteSpace(s1))
// 	//cleanUpWhiteSpace(s1)
// 	//fmt.Println(strings.Replace(s1, "(cap)", "", -1))
// 	//fmt.Println(strings.Replace(s1, "(cap)", "", -1))
// 	//re := regexp.MustCompile(`\s*\(\s*cap\s*\)\s*`)
// 	//fmt.Println(re.ReplaceAllString(s1, " "))

// }

package main

import (
	
	"go_reloaded/Functions"
	
)

func main(){
	// files := os.Args[1:]
	// //fmt.Println(inputFile[1])
	// inputFile := files[0]
	// outputFile := files[1]

	// file, inputerr := os.Open(inputFile)
	// if inputerr != nil {
	// 	fmt.Println("Error: " + inputerr.Error())
	// 	return
	// }

	// s1 := ""
	
	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	s1 += line + string('\n')
		
	// }
	// //fmt.Println(s1)

	// //s1 = go_reloaded.HexToDec(go_reloaded.BinToDec(go_reloaded.Capitalize(go_reloaded.UpperCase(go_reloaded.LowerCase(go_reloaded.GroupPunctations(go_reloaded.SinglePunctuations(go_reloaded.PairSpeechMark(go_reloaded.PunctuateVowel(s1)))))))))

	// s1 = go_reloaded.BinToDec(s1)
	// s1 = go_reloaded.HexToDec(s1)
	// s1 = go_reloaded.Capitalize(s1)
	// s1 = go_reloaded.UpperCase(s1)
	// s1 = go_reloaded.LowerCase(s1)
	// s1 = go_reloaded.GroupPunctations(s1)
	// s1 = go_reloaded.SinglePunctuations(s1)
	// s1 = go_reloaded.WordStartPunc(s1)
	// s1 = go_reloaded.PairSpeechMark(s1)
	// s1 = go_reloaded.PunctuateVowel(s1)
	// // s1 = go_reloaded.WordStartPunc(s1)


	// ofile, outputerr := os.Create(outputFile)
	// if outputerr != nil {
	// 	fmt.Println("Error: "+ outputerr.Error())
	// 	return
	// }

	// _,err := ofile.WriteString(s1)
	// if err != nil{
	// 	fmt.Println("Error: "+err.Error())
	// 	return
	// }

	//s1 := "it (cap) was the ... best of !! times, it was the ,worst (cap, 3) of times (up) , it was ' the ' a age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	// fmt.Println(go_reloaded.HexToDec("1E (hex) files were added"))
	// fmt.Println(go_reloaded.BinToDec("It has been 10 (bin) years"))
	// fmt.Println(go_reloaded.Capitalize(s1))
	// fmt.Println(go_reloaded.LowerCase(s1))
	//fmt.Println(go_reloaded.UpperCase(s1))
	//fmt.Println(go_reloaded.GroupPunctations(s1))
	//fmt.Println(go_reloaded.SinglePunctuations(s1))
	//fmt.Println(go_reloaded.PairSpeechMark(s1))
	//fmt.Println(go_reloaded.PunctuateVowel(s1))
	go_reloaded.InputToOutput()
}