package go_reloaded

import "strconv"
import "fmt"

func decimalBeforeHex(s string) string {
	words := Split(s, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			dec, err := strconv.ParseInt(words[i - 1], 16, 64)
			if err != nil{
				fmt.Println(err)
			}
			words[i-1] = strconv.Itoa(int(dec))
			words[i-1] = ToLower(words[i-1])
			//words[i] = ""
		}
	}
	return strings.Join(words, " ")
	
}