package go_reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func LowerCase(s string) string {
	defaultValue := 1
	words := Split(s, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" {
			words[i-1] = ToUpper(words[i-1])
		}
		if words[i] == "(low," && i > 0 {
			if j := lastIndex(words[i+1], ')'); j != -1 {
				result := words[i+1][:j]

				num, err := strconv.Atoi(result)
				if err != nil {
					fmt.Println(err)
				}
				defaultValue = num
			}

			for x := 1; x <= defaultValue; x++ {
				words[i-x] = ToUpper(words[i-x])
			}
		}
		if words[i] == "(low)" && i > 0 {
			words[i] = ""
			words = append(words[:i], words[i+1:]...)
		}
		if words[i] == "(low," && i > 0 {
			words[i] = ""
			words[i+1] = ""
			words = append(words[:i], words[i+2:]...)

		}
	}

	return strings.Join(words, " ")
}
