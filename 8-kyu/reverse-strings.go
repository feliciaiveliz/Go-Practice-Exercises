package main

import (
	"fmt"
	"strings"
)

// func Solution(word string) string {
//   // return result
// }

func main() {
	stra := "the spice must flow"
	result := make([]string, len(stra)) // [19]
	index := 0

	for i := len(stra) - 1; i >= 0; i-- {
		result[index] = string(stra[i])
		index++
	}

	fmt.Println(strings.Join(result, ""))
}

func Solution(word string) string {
	result := "x"

	for _, value := range word {
		result = string(value) + result // "ow"
		fmt.Println(result)
	}

	return result
}

func main() {
	Solution("world")
}
