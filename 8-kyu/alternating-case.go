package main

import (
	"fmt"
	"strings"
)

func ToAlternatingCase(str string) (result string) {
	for _, value := range str {
		if value >= 97 && value <= 122 {
			result += strings.ToUpper(string(value))
		} else if value >= 65 && value <= 90 {
			result += strings.ToLower(string(value))
		} else {
			result += string(value)
		}
	}

	fmt.Println(result)
	return
}

func main() {
	ToAlternatingCase("hello fall")
}

// func ToAlternatingCase(str string) string {
// 	result := make([]string, len(str))
//   letters := []byte(str)

// 	for idx, value := range letters {
// 		if value >= 97 && value <= 122 {
// 			result[idx] = strings.ToUpper(string(value))
// 		} else if value >= 65 && value <= 90 {
//       result[idx] = strings.ToLower(string(value))
// 		} else {
// 			result[idx] = string(value)
// 		}
// 	}

//   return strings.Join(result, "")
// }

// func main() {
// 	ToAlternatingCase("hello fall")
// }

// 97 - 122 - lower
// 65 - 90  - upper
