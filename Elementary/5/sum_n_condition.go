package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, total int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("How many numbers do you want? ")
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	fmt.Print("The sum of n-numbers given is: ", total)
}
