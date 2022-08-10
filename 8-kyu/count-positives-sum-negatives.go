package main

func CountPositivesSumNegatives(numbers []int) []int {
	count := 0
	sum := 0

	for _, value := range numbers {
		if value > 0 {
			count += 1
		} else {
			sum += value
		}
	}

	result := []int{}
	result = append(result, count, sum)
	return result
}

func main() {
	CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15})
}
