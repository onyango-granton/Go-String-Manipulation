package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	//fmt.Println(inputFile[1])
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
		s1 += line
		
	}
	fmt.Println(s1)

	ofile, outputerr := os.Create(outputFile)
	if outputerr != nil {
		fmt.Println("Error: "+ outputerr.Error())
		return
	}

	_,err := ofile.WriteString(s1)
	if err != nil{
		fmt.Println("Error: "+err.Error())
		return
	}
}