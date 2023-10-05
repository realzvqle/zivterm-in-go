package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func startApp(arg string) {
	cmd := exec.Command(arg)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("%v has not been found\n", arg)
	}
}
func startAppasAdmin(arg string) {
	cmd := exec.Command(arg)
	cmd.SysProcAttr = &syscall.SysProcAttr{Token: syscall.Token(0)}
	err := cmd.Start()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err != nil {
		fmt.Printf("%v has not been found\n", arg)
	}
}
