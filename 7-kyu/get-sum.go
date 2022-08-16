package main

// if a < b - range from a to b
// if b < a - range from b to b
// if a == b - return a

func GetSum(a, b int) (result int) {
	var smallest int
	var largest int
	if a == b {
		return a
	}

	if a < b {
		smallest = a
		largest = b
	} else if a > b {
		smallest = b
		largest = a
	}

	for n := smallest; n <= largest; n++ {
		result += n
	}

	return
}

func main() {
	GetSum(5, -1)  // 14
	GetSum(505, 4) // 127759
	GetSum(0, -1)  // -1
}
