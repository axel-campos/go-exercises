package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, total int
	fmt.Print("How many numbers do you want? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	for i := 1; i <= n; i++ {
		total += i
	}
	fmt.Print("The sum of n-numbers given is: ", total)
}
