package reverse

import "errors"

func reverse(collection []int) ([]int, error) {
	if len(collection) > 0 {
		for left, right := 0, len(collection)-1; left < right; left, right = left+1, right-1 {
			collection[left], collection[right] = collection[right], collection[left]
		}
		return collection, nil
	}
	return collection, errors.New("empty slice")
}
