package main

import (
	// "bytes"
	"fmt"
	// "io"
	"io/ioutil"
	"os"
	// "strings"

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
		// g     group
		gNode yaml.Node
	)
	fmt.Println("crossed this")
	// if err := yaml.Unmarshal(b, &g); err != nil {
	// 	return fmt.Errorf("UNMARSHAL_ERROR: %s", err.Error())
	// }

	if err := yaml.Unmarshal(b, &gNode); err != nil {
		return fmt.Errorf("UNMARSHAL_ERROR_NODE: %s", err.Error())
	}

	fmt.Println("----------------------------")
	// fmt.Println(gNode)
	// var buffer bytes.Buffer

	// addressedNode := getAddressedNode(gNode)
	// fmt.Println(addressedNode)

	// if err := yaml.NewEncoder(&buffer).Encode(addressedNode); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("is this working?")
	// fmt.Println(buffer.String())

	// reWrite([]byte(buffer.String()), fileName)

	bb, err := yaml.Marshal(&gNode)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bb))

	reWrite(bb, fileName)

	return nil
}

func getAddressedNode(n groupNode) (inst groupNodePtr) {

	var ptr []*rulesNodePtr
	for _, r := range n.Rules {
		tmp := rulesNodePtr{
			Name: &yaml.Node{
				Kind:        yaml.ScalarNode,
				LineComment: r.Name.LineComment,
				HeadComment: r.Name.HeadComment,
				FootComment: r.Name.FootComment,
				Value:       r.Name.Value,
			},
			Expression: &yaml.Node{
				Kind:        yaml.ScalarNode,
				LineComment: r.Expression.LineComment,
				HeadComment: r.Expression.HeadComment,
				FootComment: r.Expression.FootComment,
				Value:       r.Expression.Value,
			},
		}

		ptr = append(ptr, &tmp)
	}

	fmt.Println("details below")
	fmt.Println(n.Name.HeadComment)
	fmt.Println(n.Name.LineComment)
	fmt.Println(n.Name.FootComment)
	inst = groupNodePtr{
		Name: &yaml.Node{
			Kind:        yaml.ScalarNode,
			LineComment: n.Name.LineComment,
			HeadComment: n.Name.HeadComment,
			FootComment: n.Name.FootComment,
			Value:       n.Name.Value,
		},
		Rules: ptr,
	}

	return
}

// reWrite flushes the group content into the file
// thereby maintaining the recommended YAML syntax structure.
func reWrite(b []byte, fileName string) error {

	if err := ioutil.WriteFile(fileName, b, 0777); err != nil {
		return err
	}

	return nil
}
