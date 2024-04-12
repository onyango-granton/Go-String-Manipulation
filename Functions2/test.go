package go_reloaded

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	//"unicode"
)

func hexToDec(hexStr string) string {
	decValue, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return hexStr
	}
	return strconv.FormatInt(decValue, 10)
}

func binToDec(binStr string) string {
	decValue, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return binStr
	}
	return strconv.FormatInt(decValue, 10)
}

func processText(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		switch {
		case strings.Contains(word, "(hex)"):
			if i > 0 {
				words[i-1] = hexToDec(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index after removal
		case strings.Contains(word, "(bin)"):
			if i > 0 {
				words[i-1] = binToDec(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index after removal
		case strings.Contains(word, "(up)"):
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index after removal
		// Implement other rules (low, cap, punctuation, ' ,a) here
		}
	}
	return strings.Join(words, " ")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var processedLines []string

	for scanner.Scan() {
		line := scanner.Text()
		processedLine := processText(line)
		processedLines = append(processedLines, processedLine)
	}

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for _, line := range processedLines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}