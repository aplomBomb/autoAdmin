package main

import (
	"fmt"
	"os/exec"
)

func main() {
	err := exec.Command("cmd.exe", "/C", "start", "E:\\peckerheadCity\\Master").Run()
	if err != nil {
		fmt.Println(err)
	}
}
