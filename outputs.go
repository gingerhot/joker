// Copyright 2020 B1nj0y <idegorepl@gmail.com>

package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

type Output struct {
	Cmd      string
	Name     string
	Expected string
}

func CheckOutput(o Output) string {
	return outputPrinter(o, runCmd(o.Cmd) == o.Expected)
}

func runCmd(cmd string) string {
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "| impossible output <|> 不可能出现的结果 |"
	}
	return strings.TrimSpace(string(output))
}

func outputPrinter(o Output, b bool) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	info := map[bool]string{true: green("yes"), false: red("not")}[b]
	return fmt.Sprintf("Check command result %s(%s) match: %s", cyan(o.Name), cyan(o.Cmd), info)
}
