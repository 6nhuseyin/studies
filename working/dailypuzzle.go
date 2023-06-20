package main

import (
	"fmt"
	"os"
	"strings"
)

func Check(e error) {
	if e!= nil {
		panic(e)
	}
}


func controller(n []string) []string {
	var result []string
	
	return result
}

func main() {
	fileName := os.Args[1]
	readedText, err := os.ReadFile(fileName)
	Check(err)
	//fmt.Println(readedText)
	kelimeler := strings.Fields(string(readedText))
	fmt.Println(kelimeler)
	sonuc := strings.Join(kelimeler, ", ")

	aaa := []byte(sonuc)
	err = os.WriteFile(os.Args[2], aaa, 0644)
	Check(err)
}