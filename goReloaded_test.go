package go_reloaded

import (
	"fmt"
	go_reloaded "go_reloaded/Functions"
	"os"
	"testing"
)

func TestFunc(t *testing.T) {
	testFile := "Main/sample.txt"

	data, err := os.ReadFile(testFile)
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	msg,err2 := go_reloaded.OutputFunc(string(data))
	fmt.Println(msg)

	if err2 != nil {
		t.Fatalf(err2.Error())
	}

}