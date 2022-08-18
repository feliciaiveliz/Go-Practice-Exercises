package main

import "fmt"

// loop from number down to 0 -- break
// if number is a multiple of 3 || 5
// result += number
// return number

func Multiple3And5(number int) int {
	result := 0

	for i := number - 1; i > 0; i-- {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}

	fmt.Println(result)
	return result
}

func main() {
	Multiple3And5(10) // 23
}
