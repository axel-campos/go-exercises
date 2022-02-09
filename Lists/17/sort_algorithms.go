package seventeen

func SelectionSort(coll *[]int) {
	for i := 1; i < len(*coll); i++ {
		for j := i; j > 0 && (*coll)[j-1] > (*coll)[j]; j-- {
			(*coll)[j], (*coll)[j-1] = (*coll)[j-1], (*coll)[j]
		}
	}
}

func InsertionSort(coll *[]int) {
	var min int
	for i := 0; i < len(*coll)-1; i++ {
		min = i
		for j := i + 1; j < len(*coll); j++ {
			if (*coll)[j] < (*coll)[min] {
				min = j
			}
		}
		if min != i {
			(*coll)[i], (*coll)[min] = (*coll)[min], (*coll)[i]
		}
	}
}

func MergeSort(coll []int) []int {

	if len(coll) <= 1 {
		return coll
	}

	var length int = len(coll) / 2
	left := (coll)[:length]
	right := (coll)[length:]

	left = MergeSort(left)
	right = MergeSort(right)

	println("LEFT: ", left)
	println("RIGHT: ", right)
	return merge(left, right)
}

func merge(coll []int, coll2 []int) []int {
	var result []int
	for len(coll) > 0 && len(coll2) > 0 {
		if coll[0] <= coll2[0] {
			result = append(result, coll[0])
			coll = coll[1:]
		} else {
			result = append(result, coll2[0])
			coll2 = coll2[1:]
		}
	}
	for len(coll) > 0 {
		result = append(result, coll[0])
		coll = coll[1:]
	}
	for len(coll2) > 0 {
		result = append(result, coll2[0])
		coll2 = coll2[1:]
	}
	println("ORDERED: ", result)
	return result
}

func QuickSort(coll *[]int, low int, hgh int) {
	if low >= hgh || low < 0 {
		return
	}
	p := partition(coll, low, hgh)
	QuickSort(coll, low, p-1)
	QuickSort(coll, p+1, hgh)
}

func partition(coll *[]int, low int, hgh int) int {
	pivot := (*coll)[hgh]
	i := low - 1
	for j := low; j < hgh; j++ {
		if (*coll)[j] <= pivot {
			i += 1
			(*coll)[i], (*coll)[j] = (*coll)[j], (*coll)[i]
		}

	}
	i += 1
	(*coll)[i], (*coll)[hgh] = (*coll)[hgh], (*coll)[i]
	return i
}

func StoogeSort(coll *[]int, low int, hgh int) {
	if (*coll)[low] > (*coll)[hgh] {
		(*coll)[low], (*coll)[hgh] = (*coll)[hgh], (*coll)[low]
	}
	if (hgh - low + 1) > 2 {
		var pivot int = (hgh - low + 1) / 3
		StoogeSort(coll, low, hgh-pivot)
		StoogeSort(coll, low+pivot, hgh)
		StoogeSort(coll, low, hgh-pivot)
	}
}
