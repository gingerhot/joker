// Copyright 2020 B1nj0y <idegorepl@gmail.com>

package main

import (
	"fmt"
	"net"
)

type Port struct {
	Port int
	Name string
}

// Check if the port is serverd
func CheckPort(p Port) {
	portPrinter(p, portIsServed(p.Port))
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

func portPrinter(p Port, b bool) {
	result := map[bool]string{true: green("yes"), false: red("not")}[b]
	fmt.Printf("Check port %s(:%s) open  ...  %s\n", cyan(p.Name), cyan(p.Port), result)
}
