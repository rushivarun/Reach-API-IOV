package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func execute() {

	s := "iwlist scanning"
	arg := strings.Split(s, " ")
	cmd := exec.Command(arg[0], arg[1:]...)
	fmt.Println(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error : ", err)
	}

	fmt.Println("command successfully executed...")
	output := string(out)

	fmt.Println(output)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Command can only be run on a linux machine")
	} else {
		execute()
	}
}
