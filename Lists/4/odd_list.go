package four

import "errors"

func Odd_elements(collection []int) ([]int, error) {
	if len(collection) > 0 {
		var odd_collection []int
		for i := 0; i < len(collection); i = i + 2 {
			odd_collection = append(odd_collection, collection[i])
		}
		return odd_collection, nil
	}
	return nil, errors.New("empty slice")
}
