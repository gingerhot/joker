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

func CheckCmd(c Cmd, v bool) {
	if c.Cmd == "" {
		cmdPrinter(c, false, false)
	}
	cmdPrinter(c, runChech(c.Cmd), v)
}

func runChech(cmd string) bool {
	c := exec.Command("bash", "-c", "type "+cmd+" >/dev/null 2>&1")
	err := c.Run()
	if err != nil {
		return false
	}
	return true
}

func cmdPrinter(c Cmd, ok, verbose bool) {
	result := map[bool]string{true: green("yes"), false: red("not")}[ok]
	if verbose {
		fmt.Printf("Check command %s(%s) exist  ...  %s\n", cyan(c.Name), cyan(c.Cmd), result)
	} else {
		fmt.Printf("Check command %s exist  ...  %s\n", cyan(c.Name), result)
	}
}
