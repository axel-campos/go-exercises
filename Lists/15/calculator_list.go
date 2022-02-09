package fifteen

func AddLists(coll []int, coll2 []int) []int {
	var resto, temp_operation int
	var result []int
	for i, j := len(coll)-1, len(coll2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		if j < 0 {
			temp_operation = resto + coll[i]
		} else if i < 0 {
			temp_operation = resto + coll2[j]
		} else {
			temp_operation = resto + coll[i] + coll2[j]
		}
		resto = temp_operation / 10
		result = append([]int{temp_operation % 10}, result...)
	}
	if resto > 0 {
		result = append([]int{resto}, result...)
	}
	return result
}

func SubstractLists(coll []int, coll2 []int) []int {
	var temp_operation int
	var result []int
	var len_coll1, len_coll2 int = len(coll), len(coll2)
	var isNegative, isComplement bool = false, false
	if len_coll1 < len_coll2 {
		coll, coll2 = coll2, coll
		isNegative = true
	}
	for i, j := 0, len_coll2-len_coll1; i < len_coll1; i, j = i+1, j+1 {
		if j < 0 {
			temp_operation = coll[i]
		} else {
			if coll2[j] > coll[i] {
				temp_operation = 10 + coll[i] - coll2[j]
				for k := len(result) - 1; ; k-- {
					if k < 0 || (k == 0 && result[k] == 0) {
						isNegative = true
						isComplement = true
						break
					}
					if result[k] == 0 {
						result[k] = 9
					} else {
						result[k] -= 1
						break
					}
				}
			} else {
				temp_operation = coll[i] - coll2[j]
			}
		}
		result = append(result, temp_operation)
	}
	trimList(&result)
	if isComplement {
		complement_list := make([]int, len(result)+1)
		complement_list[0] = 1
		for i := 1; i < len(complement_list)-1; i++ {
			complement_list[i] = 0
		}
		result = SubstractLists(complement_list, result)
	}
	if isNegative {
		result[0] = -result[0]
	}
	return result
}

func MultiplyListsKaratsuba(coll []int, coll2 []int) []int {
	if len(coll) < 3 || len(coll2) < 3 {
		return MultiplyLists(coll, coll2)
	}
	var m int = minListBase10(coll, coll2)
	var m2 int = m / 2

	high1, low1 := spltLists(coll, m2)
	high2, low2 := spltLists(coll2, m2)

	z0 := MultiplyListsKaratsuba(low1, low2)
	z1 := MultiplyListsKaratsuba(AddLists(low1, high1), AddLists(low2, high2))
	z2 := MultiplyListsKaratsuba(high1, high2)

	xm2 := make([]int, m2*2+1)
	xm2[0] = 1
	xm1 := make([]int, m2+1)
	xm1[0] = 1

	x1 := MultiplyLists(z2, xm2)
	x2 := MultiplyLists(SubstractLists(SubstractLists(z1, z2), z0), xm1)
	return AddLists(AddLists(x1, x2), z0)
}

func MultiplyLists(coll []int, coll2 []int) []int {
	var resto int
	result := make([]int, len(coll)+len(coll2))
	for i := len(coll2) - 1; i >= 0; i-- {
		resto = 0
		for j := len(coll) - 1; j >= 0; j-- {
			result[i+j+1] += (coll2[i] * coll[j]) + resto
			resto = result[i+j+1] / 10
			result[i+j+1] %= 10
		}
		result[i] = resto
	}
	trimList(&result)
	return result
}

func trimList(coll *[]int) {
	if len(*coll) > 0 && (*coll)[0] == 0 {
		for i := 0; (*coll)[i] == 0 && len(*coll) != 1; {
			*coll = (*coll)[1:]
		}
	}
}

func spltLists(coll []int, m int) ([]int, []int) {
	m = len(coll) - m
	return coll[:m], coll[m:]
}

func minListBase10(coll []int, coll2 []int) int {
	if len(coll) > len(coll2) {
		return len(coll2)
	} else {
		return len(coll)
	}
}

/* func printList(coll []int) {
	print("[")
	for _, x := range coll {
		print(x, " ")
	}
	println("]")
} */
