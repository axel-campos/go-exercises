package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 10
	randomNumber := rand.Intn(max-min+1) + min

	scanner := bufio.NewScanner(os.Stdin)
	inputNumber := min - 1
	attempts := 0
	numbersTried := make(map[int]bool)

	for inputNumber != randomNumber {
		fmt.Print("Guess number between ", min, " and ", max, " : ")
		scanner.Scan()
		inputNumber, _ = strconv.Atoi(scanner.Text())
		if inputNumber == randomNumber {
			fmt.Println("You guess the number! Congratulations! Attempts: ", attempts)
		} else {
			_, isNumberTried := numbersTried[inputNumber]
			if !isNumberTried {
				numbersTried[inputNumber] = true
				attempts++
			}
			if inputNumber > randomNumber {
				fmt.Println("Wrong! It's a minor number. Attempts: ", attempts)
			} else {
				fmt.Println("Wrong! It's a major number. Attempts: ", attempts)
			}
		}
	}

}
