package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("O namae wa? ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("Konnichiwa,", text)
}
