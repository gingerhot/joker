// Copyright 2020 B1nj0y <idegorepl@gmail.com>

package main

import (
	"fmt"
	"os"
)

type Env struct {
	Name     string
	Env      string
	Expected string
}

func CheckEnv(e Env) {
	envPrinter(e, os.Getenv(e.Env) == e.Expected)
}

func envPrinter(e Env, b bool) {
	result := map[bool]string{true: green("yes"), false: red("not")}[b]
	fmt.Printf("Check env varible %s(%s=%s) match  ...  %s\n", cyan(e.Name), cyan(e.Env), e.Expected, result)
}
