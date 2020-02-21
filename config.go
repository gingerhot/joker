package main

import (
	"io/ioutil"
	"log"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

// colors for printer
var (
	cyan  = color.New(color.FgCyan).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

type Config struct {
	Ports   []Port
	Cmds    []Cmd
	Paths   []Path
	Envs    []Env
	Outputs []Output
}

func ReadConfig(f string) Config {
	conf := Config{}
	yamlContent, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("read config file error: %v", err)
	}
	err = yaml.Unmarshal(yamlContent, &conf)
	if err != nil {
		log.Fatalf("parse config error: %v", err)
	}
	return conf
}
