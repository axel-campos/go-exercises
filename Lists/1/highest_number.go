package highest_number

import "errors"

func highest_number(collection []int) (int, error) {
	var highest_number int = 0
	if len(collection) > 0 {
		for _, number := range collection {
			if highest_number < number {
				highest_number = number
			}
		}
		return highest_number, nil
	}
	return 0, errors.New("empty slice")
}
