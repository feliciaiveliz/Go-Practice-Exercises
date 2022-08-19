package main

import (
	"fmt"
	"strings"
)

// return count of duplicates
// case insensitive

// create a map of letters and their counts
// iterate over map and check if any letter occurs more than once, increase 'count' each time if so

func duplicate_count(s1 string) (result int) {
	letterCounts := make(map[string]int) 

	for _, letter := range s1 {
		character := strings.ToLower(string(letter))
    letterCounts[character] += 1
	}

	for _, count := range letterCounts {
    if count > 1 { result++ }
	}

	fmt.Println(result)
	return   
}

func main() {
	duplicate_count("abcde")          // 0 
	duplicate_count("abcdea")         // 1
	duplicate_count("abcdeaB11")      // 3
	duplicate_count("indivisibility") // 1
}
