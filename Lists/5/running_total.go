package five

import (
	"errors"
)

func Running_total(collection []int) (int, error) {
	if len(collection) > 0 {
		var sum int = 0
		for _, a := range collection {
			sum = sum + a
		}
		return sum, nil
	}
	return 0, errors.New("empty slice")
}
