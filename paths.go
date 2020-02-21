// Copyright 2020 B1nj0y <idegorepl@gmail.com>

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Path struct {
	Path string
	Name string
}

func CheckPath(p Path) {
	pathPrinter(p, doCheck(p.Path))
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

func pathPrinter(p Path, b bool) {
	result := map[bool]string{true: green("yes"), false: red("not")}[b]
	fmt.Printf("Check path %s(%s) exist  ...  %s\n", cyan(p.Name), cyan(p.Path), result)
}
