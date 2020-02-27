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

func CheckOutput(o Output, verbose bool) {
	outputPrinter(o, runCmd(o.Cmd) == o.Expected, verbose)
}

func runCmd(cmd string) string {
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "| impossible output <|> 不可能出现的结果 |"
	}
	return strings.TrimSpace(string(output))
}

func outputPrinter(o Output, ok, verbose bool) {
	result := map[bool]string{true: green("yes"), false: red("not")}[ok]
	if verbose {
		fmt.Printf("Check command result %s(%s) match  ...  %s\n", cyan(o.Name), cyan(o.Cmd), result)
	} else {
		fmt.Printf("Check command result %s match  ...  %s\n", cyan(o.Name), result)
	}
}
