// Copyright 2020 B1nj0y <idegorepl@gmail.com>

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type Path struct {
	Path string
	Name string
}

func CheckPath(p Path) string {
	return pathPrinter(p, doCheck(p.Path))
}

func doCheck(path string) bool {
	if path == "" {
		return false
	}
	// for relative paths
	if !strings.HasPrefix(path, "/") {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(pwd, path)
	}
	_, err := os.Stat(string(path))
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func pathPrinter(p Path, b bool) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	info := map[bool]string{true: green("yes"), false: red("not")}[b]
	return fmt.Sprintf("Check path %s(%s) exist: %s", cyan(p.Name), cyan(p.Path), info)
}
