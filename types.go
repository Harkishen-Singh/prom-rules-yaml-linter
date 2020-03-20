package main

import (
	yaml "gopkg.in/yaml.v3"
)

type rules struct {
	Name       string `yaml:"name"`
	Expression string `yaml:"expr"`
}

type rulesNode struct {
	Name       yaml.Node `yaml:"name"`
	Expression yaml.Node `yaml:"expression"`
}

type group struct {
	Name  string  `yaml:"name"`
	Rules []rules `yaml:"rules"`
}

type groupNode struct {
	Name  yaml.Node   `yaml:"name"`
	Rules []rulesNode `yaml:"rules"`
}
