package seven

import "errors"

func SumFor(collection []int) (int, error) {
	if len(collection) > 0 {
		var sum int
		for _, a := range collection {
			sum = sum + a
		}
		return sum, nil
	}
	return 0, errors.New("empty slice")
}

func SumRecursion(collection []int) (int, error) {
	if len(collection) > 0 {
		sum, _ := SumRecursion(collection[1:])
		return collection[0] + sum, nil
	}
	return 0, errors.New("empty slice")
}
