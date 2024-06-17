package main

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	data := []map[string]interface{}{
		{"name": "Elliot",
			"age": 25},
	}
	filePath := "/home/george/goproject/src/student.vkusvill.ru/George/go-course/golang_study_ak/course1/13.popular_package/9.package_yaml/task1.13.9.4/test/file.yaml"
	err := writeYAML(filePath, data)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func writeYAML(filePath string, data interface{}) error {
	path := filepath.Dir(filePath)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	infoToWrite, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, []byte(infoToWrite), 0644)
	return err
}
