package main

import (
	"fmt"
	// "strings"
)

func HasUniqueChar(str string) bool {
	obj := make(map[string]int)

	for i := 0; i < len(str); i++ {
		// value, boolean := obj[string(str[i])]
		fmt.Println(string(str[i]))
	}

	fmt.Println(obj)
	return true
}

func main() {
	HasUniqueChar("  nAa")
}
