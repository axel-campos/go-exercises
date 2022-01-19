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
	fmt.Print("Which operator do you want? + or *? ")
	scanner.Scan()
	operator := scanner.Text()
	if operator == "+" {
		total = 0
		for i := 1; i <= n; i++ {
			total += i
		}
	} else if operator == "*" {
		total = 1
		for i := 1; i <= n; i++ {
			total *= i
		}
	} else {
		fmt.Print("Not a valid option")
		return
	}
	fmt.Print("The result of n-numbers given is: ", total)
}
