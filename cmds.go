// Copyright 2020 B1nj0y <idegorepl@gmail.com>

package main

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

type Cmd struct {
	Cmd  string
	Name string
}

func CheckCmd(c Cmd) string {
	return cmdPrinter(c, runChech(c.Cmd))
}

func runChech(cmd string) bool {
	c := exec.Command("type", cmd, ">", "/dev/null", "2>&1")
	err := c.Run()
	if err != nil {
		return false
	}
	return true
}

func cmdPrinter(c Cmd, b bool) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	info := map[bool]string{true: green("yes"), false: red("not")}[b]
	return fmt.Sprintf("Check command %s(%s) exist: %s", cyan(c.Name), cyan(c.Cmd), info)
}
