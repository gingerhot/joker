// Copyright 2020 B1nj0y <idegorepl@gmail.com>

package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Env struct {
	Name     string
	Env      string
	Expected string
}

func CheckEnv(e Env) string {
	return envPrinter(e, os.Getenv(e.Env) == e.Expected)
}

func envPrinter(e Env, b bool) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	info := map[bool]string{true: green("yes"), false: red("not")}[b]
	return fmt.Sprintf("Check env varible %s(%s=%s) match: %s", cyan(e.Name), cyan(e.Env), e.Expected, info)
}
