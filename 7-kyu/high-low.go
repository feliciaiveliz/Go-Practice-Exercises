package main

import (
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(str string) (result string) {
	slice := []int{}

	for _, value := range strings.Split(str, " ") {
		n, _ := strconv.Atoi(string(value))
		slice = append(slice, n)
	}

	sort.Ints(slice)

	result += strconv.Itoa(slice[len(slice)-1])
	result += " "
	result += strconv.Itoa(slice[0])

	return result
}

func main() {
	// HighAndLow("3 1 2")
	HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")
}
