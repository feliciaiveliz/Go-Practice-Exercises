package main

import (
	"fmt"
	"sort"
)

// sort numbers
// triangle:
// a == b == c
// a != b && b == c
// a - b == 1 / b - c == 1 / a - c == 1

func IsTriangle(a, b, c int) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println(false)
		return false
	}
	if (a + b + c) <= 0 {
		fmt.Println(false)
		return false
	}

	sides := []int{a, b, c}
	sort.Ints(sides)
	fmt.Println(sides)

	a = sides[0]
	b = sides[1]
	c = sides[2]

	if (a != b && b == c) ||
		(a == b && b == c && c == a) ||
		(a+1 == b && b+1 == c && a+2 == c) {
		fmt.Println(true)
		return true
	}

	fmt.Println(false)
	return false
}

func main() {
	IsTriangle(5, 1, 2)  // false
	IsTriangle(1, 2, 5)  // false
	IsTriangle(2, 5, 1)  // false
	IsTriangle(4, 2, 3)  // true
	IsTriangle(5, 1, 5)  // true
	IsTriangle(2, 2, 2)  // true
	IsTriangle(-1, 2, 3) // false
}
