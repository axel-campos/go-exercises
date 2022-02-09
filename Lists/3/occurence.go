package three

import "errors"

func Ocurrence(collection []int, number int) (bool, error) {
	if len(collection) > 0 {
		for _, a := range collection {
			if number == a {
				return true, nil
			}
		}
		return false, nil
	}
	return false, errors.New("empty slice")

}
