package main

func Feast(beast string, dish string) bool {
	bFirst, bLast := string(beast[0]), string(beast[len(beast)-1])
	dFirst, dLast := string(dish[0]), string(dish[len(dish)-1])

	if bFirst == dFirst && bLast == dLast {
		return true
	}

	return false
}

func main() {
	Feast("great blue heron", "garlic naan")
}
