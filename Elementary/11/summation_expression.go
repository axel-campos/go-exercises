package main

import "math"

func main() {
	k := 1.00
	var sum float64 = 0
	for k <= math.Pow(10, 6) {
		sum += (math.Pow(-1, k+1)) / (2*k - 1)
		k++
	}
	result := 4 * sum
	println("The resut of summation is: ", result)
}
