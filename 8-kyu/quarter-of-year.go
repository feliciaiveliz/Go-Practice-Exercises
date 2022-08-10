package main

func QuarterOf(month int) int {
	switch month {
	case 1, 2, 3:
		return 1
	case 4, 5, 6:
		return 2
	case 7, 8, 9:
		return 3
	default:
		return 4
	}
}

func QuarterOf(month int) int {
	return month + 2/3
}
