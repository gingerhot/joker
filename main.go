package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	help := flag.Bool("h", false, "help")
	verbose := flag.Bool("v", false, "vorbose")
	flag.Parse()

	if len(flag.Args()) == 0 && *help {
		usage()
		os.Exit(0)
	}

	confPath := "joker.yaml"
	if len(flag.Args()) == 1 {
		confPath = flag.Args()[0]
	}

	conf := ReadConfig(confPath)

	for _, port := range conf.Ports {
		CheckPort(port)
	}

	for _, cmd := range conf.Cmds {
		CheckCmd(cmd, *verbose)
	}

	for _, env := range conf.Envs {
		CheckEnv(env, *verbose)
	}

	for _, output := range conf.Outputs {
		CheckOutput(output, *verbose)
	}

	for _, path := range conf.Paths {
		CheckPath(path, *verbose)
	}
}

func usage() {
	helpMsg := `
Joker - A commandline tool helps to check your dev environment

Usage: joker [options...] <a config yaml>

Just run joker if your config file named exactly "joker.yaml"

-h    Help
-v    Verbose mode, show checking item in detail

`
	fmt.Println(helpMsg)
}
