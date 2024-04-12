package go_reloaded

import "os"
import "fmt"
import "bufio"

func InputToOutput(){
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
		s1 += line + string('\n')
		
	}
	//fmt.Println(s1)

	//s1 = go_reloaded.HexToDec(go_reloaded.BinToDec(go_reloaded.Capitalize(go_reloaded.UpperCase(go_reloaded.LowerCase(go_reloaded.GroupPunctations(go_reloaded.SinglePunctuations(go_reloaded.PairSpeechMark(go_reloaded.PunctuateVowel(s1)))))))))

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
	// s1 = go_reloaded.WordStartPunc(s1)

	s2,err1 := OutputFunc(s1)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}


	ofile, outputerr := os.Create(outputFile)
	if outputerr != nil {
		fmt.Println("Error: "+ outputerr.Error())
		return
	}

	_,err := ofile.WriteString(s2)
	if err != nil{
		fmt.Println("Error: "+err.Error())
		return
	}
}