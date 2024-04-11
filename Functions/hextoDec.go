package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := processLine(line)
		_, err = writer.WriteString(modifiedLine + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing output file:", err)
		return
	}
}

func processLine(line string) string {
	var modifiedLine strings.Builder

	words := strings.Split(line, " ")
	for i, word := range words {
		modifiedWord := word

		if strings.HasSuffix(word, "(hex)") {
			hexNum := strings.TrimSuffix(word, "(hex)")
			decNum, _ := strconv.ParseInt(hexNum, 16, 64)
			modifiedWord = strconv.FormatInt(decNum, 10)
		} else if strings.HasSuffix(word, "(bin)") {
			binNum := strings.TrimSuffix(word, "(bin)")
			decNum, _ := strconv.ParseInt(binNum, 2, 64)
			modifiedWord = strconv.FormatInt(decNum, 10)
		} else if strings.HasSuffix(word, "(up)") {
			modifiedWord = strings.TrimSuffix(word, "(up)")
			modifiedWord = strings.ToUpper(modifiedWord)
		} else if strings.HasSuffix(word, "(low)") {
			modifiedWord = strings.TrimSuffix(word, "(low)")
			modifiedWord = strings.ToLower(modifiedWord)
		} else if strings.HasSuffix(word, "(cap)") {
			modifiedWord = strings.TrimSuffix(word, "(cap)")
			modifiedWord = strings.Title(modifiedWord)
		} else if strings.Contains(word, "(low, ") {
			parts := strings.Split(word, "(low, ")
			count, _ := strconv.Atoi(strings.TrimSuffix(parts[1], ")"))
			for i, w := range strings.Fields(parts[0]) {
				if i < count {
					modifiedWord += strings.ToLower(w) + " "
				} else {
					modifiedWord += w + " "
				}
			}
			modifiedWord = strings.TrimSuffix(modifiedWord, " ")
		} else if strings.Contains(word, "(up, ") {
			parts := strings.Split(word, "(up, ")
			count, _ := strconv.Atoi(strings.TrimSuffix(parts[1], ")"))
			for i, w := range strings.Fields(parts[0]) {
				if i < count {
					modifiedWord += strings.ToUpper(w) + " "
				} else {
					modifiedWord += w + " "
				}
			}
			modifiedWord = strings.TrimSuffix(modifiedWord, " ")
		} else if i < len(words)-1 && strings.HasPrefix(words[i+1], "(cap,") {
			modifiedWord = strings.Title(word)
			count, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(words[i+1], "(cap,"), ")"))
			for j := i + 2; j < i+2+count && j < len(words); j++ {
				modifiedWord += " " + strings.Title(words[j])
			}
			i += 1 + count
		}

		if strings.ContainsAny(modifiedWord, ".,!?:;") {
			modifiedWord = formatPunctuation(modifiedWord)
		}

		if strings.Contains(modifiedWord, "'") {
			modifiedWord = formatQuotes(modifiedWord)
		}

		if strings.HasPrefix(modifiedWord, "a ") && isVowel([]rune(modifiedWord)[2]) {
			modifiedWord = "an " + modifiedWord[2:]
		}

		modifiedLine.WriteString(modifiedWord + " ")
	}

	return strings.TrimSuffix(modifiedLine.String(), " ")
}

func formatPunctuation(word string) string {
	var modifiedWord strings.Builder
	var prevRune rune
	var punctuationGroup bool

	for _, r := range word {
		if isPunctuation(r) {
			if prevRune != 0 && isPunctuation(prevRune) {
				punctuationGroup = true
			} else {
				punctuationGroup = false
				modifiedWord.WriteRune(' ')
			}
			modifiedWord.WriteRune(r)
		} else {
			if punctuationGroup {
				modifiedWord.WriteRune(r)
			} else {
				modifiedWord.WriteString(string(r))
			}
			punctuationGroup = false
		}
		prevRune = r
	}

	return modifiedWord.String()
}

func formatQuotes(word string) string {
	parts := strings.Split(word, "'")
	var modifiedWord strings.Builder

	for i, part := range parts {
		if i == 0 {
			modifiedWord.WriteString(part)
		} else if i == len(parts)-1 {
			modifiedWord.WriteString("'" + part + "'")
		} else {
			modifiedWord.WriteString("'" + part + "' ")
		}
	}

	return modifiedWord.String()
}

func isPunctuation(r rune) bool {
	return strings.ContainsRune(".,!?:;", r)
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouAEIOU", r) || r == 'h'
}