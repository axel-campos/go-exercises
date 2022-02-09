package eleven

import "errors"

func Concat_sort(collection []int, collection2 []int) ([]int, error) {
	if len(collection) != 0 || len(collection2) != 0 {
		result := make([]int, len(collection)+len(collection2))
		for i, j := 0, 0; i < len(collection)-1 && j < len(collection)-1; {
			if len(collection)-1 == i {
				result = append(result, collection[j])
				j++
			} else if len(collection2)-1 == j {
				result = append(result, collection[i])
				i++
			} else {
				if collection[i] < collection2[j] {
					result = append(result, collection[i])
					i++
				} else {
					result = append(result, collection[j])
					j++
				}
			}

		}
		return result, nil
	}
	return nil, errors.New("empty collections")
}
