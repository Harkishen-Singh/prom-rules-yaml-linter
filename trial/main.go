package main

import (
	"fmt"
	"io/ioutil"
	yaml "gopkg.in/yaml.v3"
)

type group struct {
	Name string `yaml:"name"`
	Type yaml.Node `yaml:"type"`
}

func main() {
	b, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		panic(err)
	}

	var tmp yaml.Node

	if err = yaml.Unmarshal(b, &tmp); err != nil {
		panic(err)
	}

	bb, err := yaml.Marshal(&tmp)
	if err != nil {
		panic(err)
	}

	fmt.Println(tmp)

	fmt.Println("-----------------")

	fmt.Println(string(bb))
}