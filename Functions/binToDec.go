package go_reloaded

import "strconv"
import "fmt"
import "strings"

func BinToDec(s string) string {
	words := Split(s, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			dec, err := strconv.ParseInt(words[i - 1], 2, 64)
			if err != nil{
				fmt.Println(err)
			}
			words[i-1] = strconv.Itoa(int(dec))
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
	
}