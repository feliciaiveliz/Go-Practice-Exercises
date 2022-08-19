package main

import (
	"fmt"
	"strings"
)

// Given a string of words, you need to find the highest scoring word.

// Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.

// You need to return the highest scoring word as a string.

// If two words score the same, return the word that appears earliest in the original string.

// All letters will be lowercase and all inputs will be valid.

// create a string of letters "a" to "z"
// split string on spaces " "
// iterate through slice of words, for each word:
  // split word on empty space ""
		// get score of word: get index of letter in string + 1
		// increase 'wordScore' by 'score' of current letter
  // add word and score to map scores[word] = score
	// go to next word and compute score for that word
// find the word with the highest score
// iterate over map and keep running track of the highest score
// return the word of highest score


func High(s string) (result string) {
	highestScore := 0
	scores := make(map[string]int)
	words := strings.Split(s, " ")
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, word := range words {
		currentWord := strings.Split(word, "")
		score := 0

    for _, letter := range currentWord {
			score += strings.Index(alphabet, letter) + 1
		}

    scores[word] = score
	}

	for word, score := range scores {
    if score > highestScore { 
			highestScore = score
			result = word
	  }
	}
  
  fmt.Println(result)
	return ""
}

func main() {
  High("man i need a taxi up to ubud") // taxi
	High("what time are we climbing up the volcano") // volcano
	High("aaa b") // aaa
}
