package sixteen

import (
	"errors"
	"math"
)

func ConvertBaseList(coll []int, base int, toBase int) ([]int, error) {
	if base == toBase {
		return coll, nil
	}
	for _, x := range coll {
		if x >= base {
			error := errors.New("no base according")
			return nil, error
		}
	}
	return decimalToBase(BaseToDecimal(coll, base), toBase), nil
}

func BaseToDecimal(coll []int, base int) int {
	var total int
	for i, j := 0, len(coll)-1; j >= 0; i, j = i+1, j-1 {
		total += coll[j] * int(math.Pow(float64(base), float64(i)))
	}
	return total
}

func decimalToBase(digit int, base int) []int {
	var result []int
	for digit > 0 {
		result = append([]int{digit % base}, result...)
		digit /= base
	}
	return result
}
