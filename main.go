package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func main() {
	fileList := os.Args[1:]

	for _, fileName := range fileList {
		if err := process(fileName); err != nil {
			panic(err)
		}
	}
}

// process processes the files passed as params
// which are then formatted as required.
func process(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("FILE_NOT_FOUND: %s", fileName)
	}

	var (
		gNode yaml.Node
	)

	if err := yaml.Unmarshal(b, &gNode); err != nil {
		return fmt.Errorf("UNMARSHAL_ERROR_NODE: %s", err.Error())
	}

	bb, err := yaml.Marshal(&gNode)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bb))

	reWrite(bb, fileName)

	return nil
}

// reWrite flushes the group content into the file
// thereby maintaining the recommended YAML syntax structure.
func reWrite(b []byte, fileName string) error {

	if err := ioutil.WriteFile(fileName, b, 0777); err != nil {
		return err
	}

	return nil
}
