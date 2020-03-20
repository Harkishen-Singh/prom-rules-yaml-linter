package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func main() {
	fileList := os.Args[1:]

	for _, fileName := range fileList {
		process(fileName)
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
		g     group
		gNode groupNode
	)

	if err := yaml.Unmarshal(b, &g); err != nil {
		return fmt.Errorf("UNMARSHAL_ERROR: %s", err.Error())
	}

	if err := yaml.Unmarshal(b, &gNode); err != nil {
		return fmt.Errorf("UNMARSHAL_ERROR_NODE: %s", err.Error())
	}

	reWrite(&g, fileName)

	return nil
}

// reWrite flushes the group content into the file
// thereby maintaining the recommended YAML syntax structure.
func reWrite(grp *group, fileName string) error {
	if grp != nil {
		b, err := yaml.Marshal(*grp)
		if err != nil {
			return fmt.Errorf("MARSHALL_ERROR: %v", *grp)
		}

		if err = ioutil.WriteFile(fileName, b, 0777); err != nil {
			return err
		}
	}

	return nil
}
