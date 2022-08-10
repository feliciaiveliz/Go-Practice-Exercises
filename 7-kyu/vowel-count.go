package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	vowels := "aeiou"

	for _, value := range str {
		if strings.Contains(vowels, string(value)) {
			count += 1
		}
	}

	fmt.Println(count)
	return
}

func main() {
	GetCount("abracadabra")
}
