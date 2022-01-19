package main

import (
	"fmt"
)

func main() {
	i := 1
	var isPrime bool
	for {
		i++
		isPrime = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(i, "\t")
		}
	}
}
