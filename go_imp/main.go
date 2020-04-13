package main

import (
	"fmt"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

func execute() {

	s := "sudo iw wlp2s0 scan"
	arg := strings.Split(s, " ")
	fmt.Println(s)
	cmd := exec.Command(arg[0], arg[1:]...)
	fmt.Println(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error is here : ", err)
	}

	fmt.Println("command successfully executed...")
	output := string(out)
	r, _ := regexp.Compile("s([a-z]+)gnal")
	fmt.Println(r.FindAllString("signal", 2))
	fmt.Println(reflect.TypeOf(output))
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Command can only be run on a linux machine")
	} else {
		execute()
	}
}
