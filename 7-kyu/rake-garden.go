package main

import (
	"fmt"
	"strings"
)

func RakeGarden(garden string) (result string) {
	for _, word := range strings.Split(garden, " ") {
		if word != "rock" && word != "gravel" {
			result += "gravel" + " "
		} else {
			result += word + " "
		}
	}

	fmt.Println(result)
	return result
}

func main() {
	RakeGarden("slug spider rock gravel gravel gravel gravel gravel gravel gravel")
}
