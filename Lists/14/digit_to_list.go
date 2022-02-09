package fourteen

func DigitToList(digit int) []int {
	var result []int
	for digit > 0 {
		result = append([]int{digit % 10}, result...)
		digit /= 10
	}
	return result
}
