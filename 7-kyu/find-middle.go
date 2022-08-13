package main

// if odd, return 1 middle
// get length == 7
// length / 2 == 3 (round up)
// return string[middle]
// if even, return 2 middle
// get length == 4
// length / 2 == 2
// slice from 1 - 2
// string[middle - 1: middle]

func GetMiddle(s string) string {
	middle := len(s) / 2

	if len(s)%2 == 1 {
		return string(s[middle])
	} else {
		return s[middle-1 : middle+1]
	}
}

func main() {
	GetMiddle("testing") // "t"
	GetMiddle("test")    // "es"
}
