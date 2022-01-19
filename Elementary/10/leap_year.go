package main

import (
	"fmt"
	"time"
)

func main() {
	actualYear := time.Now().Year()
	ocurrences := 0
	for ocurrences < 20 {
		if actualYear%4 == 0 && (actualYear%100 != 0 || actualYear%400 == 0) {
			fmt.Println(actualYear)
			ocurrences++
		}
		actualYear++
	}
}
