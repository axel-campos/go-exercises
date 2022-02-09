package six

import "errors"

func IsPalindrome(test string) (bool, error) {
	if len(test) > 0 {
		for left, right := 0, len(test)-1; left < right; left, right = left+1, right-1 {
			if test[left] != test[right] {
				return false, nil
			}
		}
		return true, nil
	}
	return false, errors.New("empty string")
}
