package eight

func ToSquareAll(collection []int) ([]int, error) {
	return Collect(collection, ToSquare), nil
}

func Collect(collection []int, f func(int) int) []int {
	result := make([]int, len(collection))
	for i, item := range collection {
		result[i] = f(item)
	}
	return result
}

func ToSquare(value int) int {
	return value * value
}
