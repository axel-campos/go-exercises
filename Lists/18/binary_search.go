package eighteen

func BinarySearch(coll []int, n int, T int) bool {
	L := 0
	R := n - 1
	for L <= R {
		var m int = (L + R) / 2
		println("L,R = ", L, ",", R)
		println("m = ", m)
		if T > coll[m] {
			L = m + 1
		} else if T < coll[m] {
			R = m - 1
		} else {
			return true
		}
	}
	return false
}
