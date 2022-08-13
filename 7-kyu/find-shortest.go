package main

import (
	"sort"
	"strings"
)

func FindShort(str string) int {
	hash := make(map[string]int)
	s := strings.Fields(str)

	for _, value := range s {
		value := string(value)
		hash[value] = len(value)
	}

	wordsLengths := make([]int, len(hash))
	i := 0

	for _, value := range hash {
		wordsLengths[i] = value
		i++
	}

	sort.Ints(wordsLengths)
	return wordsLengths[0]
}

func main() {
	FindShort("Let's travel abroad shall we")                        // 2
	FindShort("bitcoin take over the world maybe who knows perhaps") // 3
}
