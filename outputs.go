// Copyright 2020 B1nj0y <idegorepl@gmail.com>

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type Output struct {
	Cmd      string
	Name     string
	Expected string
}

func CheckOutput(o Output) {
	outputPrinter(o, runCmd(o.Cmd) == o.Expected)
}

func runCmd(cmd string) string {
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "| impossible output <|> 不可能出现的结果 |"
	}
	return strings.TrimSpace(string(output))
}

func outputPrinter(o Output, b bool) {
	result := map[bool]string{true: green("yes"), false: red("not")}[b]
	fmt.Printf("Check command result %s(%s) match  ...  %s\n", cyan(o.Name), cyan(o.Cmd), result)
}
