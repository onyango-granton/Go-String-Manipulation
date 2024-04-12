package go_reloaded

import (
	"bufio"
	"fmt"
	"os"
)

func InputToOutput() {
	files := os.Args[1:]

	inputFile := files[0]
	outputFile := files[1]

	file, inputerr := os.Open(inputFile)
	if inputerr != nil {
		fmt.Println("Error: " + inputerr.Error())
		return
	}

	s1 := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		s1 += line + string('\n')

	}

	s2, err1 := OutputFunc(s1)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}

	ofile, outputerr := os.Create(outputFile)
	if outputerr != nil {
		fmt.Println("Error: " + outputerr.Error())
		return
	}

	_, err := ofile.WriteString(s2)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
}
