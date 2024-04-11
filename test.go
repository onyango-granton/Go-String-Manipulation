package main

import (
	"fmt"
	"regexp"
	//"strings"
)

func formatPunctuation(text string) string {
	re := regexp.MustCompile(`(\w)([.,!?;:])(\s+)([.,!?;:])(\s+)([.,!?;:])`)
	text = re.ReplaceAllString(text, "${1}${2}${4}${6}")

	re = regexp.MustCompile(`(\w)([.,!?;:])(\s+)([.,!?;:])`)
	text = re.ReplaceAllString(text, "${1}${2}${4}")

	return text
}

func main() {
	text := "I was sitting over there ,and then BAMM !!"
	formattedText := formatPunctuation(text)
	fmt.Println(formattedText)

	text = "I was thinking ... You were right"
	formattedText = formatPunctuation(text)
	fmt.Println(formattedText)
}