package main

import (
	"fmt"
	"strings"
)

func HasUniqueChar(str string) bool {
	hash := make(map[string]int)

	for _, value := range str {
		hash[string(value)] += 1
	}

	fmt.Println(hash)

	for _, value := range str {
		if hash[string(value)] > 1 {
			fmt.Println(false)
			return false
		}
	}

	fmt.Println(true)
	return true
}

func HasUniqueChar(str string) bool {
	for _, letter := range str {
		if strings.Count(str, string(letter)) > 1 {
			return false
		}
	}

	return true
}

func main() {
	HasUniqueChar("  nAa")
	HasUniqueChar("abcde")
}
