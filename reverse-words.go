package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	result := ""
	words := strings.Split(str, " ")

	for _, word := range words {
		result = word + " " + result
	}

	fmt.Println(result)
	return strings.Trim(result, " ")
}

func main() {
	ReverseWords("hello world!")
}
