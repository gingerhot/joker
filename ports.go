// Copyright 2020 B1nj0y <idegorepl@gmail.com>

package main

import (
	"fmt"
	"net"

	"github.com/fatih/color"
)

type Port struct {
	Port int
	Name string
}

// Check if the port is serverd
func CheckPort(p Port) string {
	return portPrinter(p, portIsServed(p.Port))
}

func portIsServed(port int) bool {
	// check if the port is used
	conn, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return true
	}
	conn.Close()
	// then check bind to 127.0.0.1
	conn, err = net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return true
	}
	conn.Close()
	return false
}

func portPrinter(p Port, b bool) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	info := map[bool]string{true: green("yes"), false: red("not")}[b]
	return fmt.Sprintf("Check port %s(:%s) open: %s", cyan(p.Name), cyan(p.Port), info)
}
