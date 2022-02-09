package twelve

import "errors"

func Rotate_list(collection []int, n_rotate int) ([]int, error) {
	if len(collection) > 0 {
		for i := 0; i < n_rotate; i++ {
			for j := len(collection) - 1; j < 0; j-- {
				collection[j], collection[j-1] = collection[j-1], collection[j]
			}
		}
		return collection, nil
	}
	return nil, errors.New("empty slice")
}
