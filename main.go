package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	confPath := "joker.yaml"
	if len(flag.Args()) == 1 {
		confPath = flag.Args()[0]
	}
	conf := ReadConfig(confPath)

	for _, port := range conf.Ports {
		r := CheckPort(port)
		fmt.Println(r)
	}

	for _, cmd := range conf.Cmds {
		r := CheckCmd(cmd)
		fmt.Println(r)
	}

	for _, env := range conf.Envs {
		r := CheckEnv(env)
		fmt.Println(r)
	}

	for _, output := range conf.Outputs {
		r := CheckOutput(output)
		fmt.Println(r)
	}

	for _, path := range conf.Paths {
		r := CheckPath(path)
		fmt.Println(r)
	}
}
