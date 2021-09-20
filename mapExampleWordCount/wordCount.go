package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	stringFields := strings.Fields(text)
	countWord := map[string]int{}
	
	for _,word := range stringFields {
		countWord[strings.ToLower(word)]++
	}

	fmt.Println(countWord)
}
