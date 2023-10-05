package main

import (
	"fmt"
)

func cmdChecker(cmd string, args string, ver float32) {
	switch cmd {
	case "echo":
		fmt.Println(args)
	case "start":
		startApp(args)
	case "ver":
		fmt.Printf("ZivTerm Version %0.1v\nCopyright zvqle, all rights reserved\n", ver)
	default:
		fmt.Printf("%v is not the correct syntax!\n", cmd)

	}
}
