package go_reloaded

import (
	"fmt"
	"os"
	"testing"

	go_reloaded "go_reloaded/Functions"
)

func TestFunc(t *testing.T) {
	testFile := "Main/sample.txt"

	data, err := os.ReadFile(testFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err2 := go_reloaded.OutputFunc(string(data))
	// fmt.Println(msg)

	if err2 != nil {
		t.Fatalf(err2.Error())
	}
}
