// Copyright 2020 B1nj0y <idegorepl@gmail.com>

package main

import (
	"fmt"
	"os/exec"
)

type Cmd struct {
	Cmd  string
	Name string
}

func CheckCmd(c Cmd) {
	if c.Cmd == "" {
		cmdPrinter(c, false)
	}
	cmdPrinter(c, runChech(c.Cmd))
}

func runChech(cmd string) bool {
	c := exec.Command("bash", "-c", "type "+cmd+" >/dev/null 2>&1")
	err := c.Run()
	if err != nil {
		return false
	}
	return true
}

func cmdPrinter(c Cmd, b bool) {
	result := map[bool]string{true: green("yes"), false: red("not")}[b]
	fmt.Printf("Check command %s(%s) exist  ...  %s\n", cyan(c.Name), cyan(c.Cmd), result)
}
