package go_reloaded

import "strings"
import "strconv"
import "fmt"

 func UpperCase(s string) string {
	defaultValue := 1
	words := Split(s, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)"{
			words[i-1] = ToUpper(words[i-1])
		}
		if words[i] == "(up," && i > 0 {
			if j := lastIndex(words[i + 1], ')'); j != -1{
				result := words[i+1][:j]
				//fmt.Println(result)
				num, err := strconv.Atoi(result)
				if err != nil {
					fmt.Println(err)
				}
				defaultValue = num
			}
			//fmt.Println(defaultValue)
			for x := 1; x <= defaultValue; x++ {
				words[i-x] = ToUpper(words[i-x])
			}
		}
		if words[i] == "(up)" && i > 0 {
			words[i] = ""
			words = append(words[:i], words[i+1:]...)
		}
		if words[i] == "(up," && i > 0 {
			words[i] = ""
			words[i + 1] = ""
			words = append(words[:i], words[i+2:]...)
		}
	}

	return strings.Join(words, " ")
}