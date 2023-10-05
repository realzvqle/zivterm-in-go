package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var temp string
	var err error
	var ver float32 = 0.1
	loop := 1
	scanner := bufio.NewReader(os.Stdin)
	fmt.Printf("ZivTerm Version %0.1v\nCopyright zvqle, all rights reserved\n\n", ver)

	for loop == 1 {
		fmt.Print("> ")
		temp, err = scanner.ReadString('\n')
		checkError(err)
		args := strings.Fields(strings.TrimSpace(temp))
		cmdChecker(args[0], strings.Join(args[1:], " "), ver)
	}
}
