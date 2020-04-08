package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {

	out, err := exec.Command("iw", "wlp2s0", "station", "dump", "|", "grep", "'signal avg:'").Output()

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
