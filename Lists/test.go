package main

import (
	"fmt"

	"lists.com/nineteen"
	"lists.com/twenty"
)

func main() {
	/* test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 18, 19, 20}
	fmt.Println(one.Highest_number(test))
	fmt.Println(two.Reverse(test))
	fmt.Println(three.Ocurrence(test, 30))
	fmt.Println(fourteen.DigitToList(3480))

	fmt.Println(fifteen.AddLists(fourteen.DigitToList(200), fourteen.DigitToList(300)))
	fmt.Println(fifteen.SubstractLists(fourteen.DigitToList(1000), fourteen.DigitToList(1001)))
	fmt.Println(fifteen.MultiplyLists(fourteen.DigitToList(8392), fourteen.DigitToList(8392)))
	fmt.Println(fifteen.MultiplyListsKaratsuba(fourteen.DigitToList(8392), fourteen.DigitToList(8392)))
	fmt.Println(sixteen.ConvertBaseList(fourteen.DigitToList(8392), 10, 2))
	fmt.Println(sixteen.ConvertBaseList(fourteen.DigitToList(210), 3, 10))
	fmt.Println(sixteen.ConvertBaseList([]int{10, 11, 12, 13, 14}, 16, 9))

	var test2 = []int{3, 19, 2, 16, 14, 6, 15, 7, 8, 10, 0, 18, 20, 12, 1, 4, 17, 11, 13, 5, 9}
	seventeen.SelectionSort(&test2)
	seventeen.InsertionSort(&test2)
	test2 = seventeen.MergeSort(test2)
	seventeen.QuickSort(&test2, 0, len(test2)-1)
	seventeen.StoogeSort(&test2, 0, len(test2)-1)
	fmt.Println(test2)

	var test3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16, 17, 18, 19, 20}
	fmt.Println("2= ", eighteen.BinarySearch(test3, len(test3), 18)) */

	var test4 = []string{"Hello", "World", "in", "a", "frame"}
	nineteen.PrintFrame(test4)

	var test5 string = "The quick brown fox"
	fmt.Println(twenty.NormalToPigLatin(test5))
}
