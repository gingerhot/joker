package main

import "flag"

func main() {
	flag.Parse()
	confPath := "joker.yaml"
	if len(flag.Args()) == 1 {
		confPath = flag.Args()[0]
	}
	conf := ReadConfig(confPath)

	for _, port := range conf.Ports {
		CheckPort(port)
	}

	for _, cmd := range conf.Cmds {
		CheckCmd(cmd)
	}

	for _, env := range conf.Envs {
		CheckEnv(env)
	}

	for _, output := range conf.Outputs {
		CheckOutput(output)
	}

	for _, path := range conf.Paths {
		CheckPath(path)
	}
}
