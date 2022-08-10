package main

import "fmt"

func ClosestMultipleOf10(n uint32) uint32 {
	lower := n
	higher := n

	for lower > 0 || higher < 1000 {
		lower -= 1
		higher += 1

		if n%10 == 0 {
			return n
		} else if higher%10 == 0 {
			fmt.Println(higher)
			return higher
		} else if lower%10 == 0 {
			fmt.Println(lower)
			return lower
		}
	}

	return 0
}

func main() {
	ClosestMultipleOf10(1964030160) // 1964030160
}
