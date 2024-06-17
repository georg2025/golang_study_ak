package main

import (
	"fmt"

	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

func main() {
}

type Person struct {
	Name string `json:"name" yaml: "name"`
	Age  int    `json: "age" yaml: "age"`
}

func unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err == nil {
		return err
	}
	err = yaml.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("cant decode neither JSON, nor YAML")
	}
	return nil
}
