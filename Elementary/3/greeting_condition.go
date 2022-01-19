package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("O namae wa? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	if strings.ToUpper(name) == "ALICE" || strings.ToUpper(name) == "BOB" {
		fmt.Println("Konnichiwa,", name)
	} else {
		fmt.Println("Anata no namae wa Bob to Alice de wa arimasen. Sayounara")
	}

}
