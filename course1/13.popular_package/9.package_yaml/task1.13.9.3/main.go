package main

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	comment1 := Comment{Text: "hello"}
	comment2 := Comment{Text: "go"}
	comment3 := Comment{Text: "moo"}
	comment4 := Comment{Text: "what"}
	mary := User{Name: "Mary", Age: 32, Comments: []Comment{comment1, comment2}}
	ivan := User{Name: "Ivan", Age: 21, Comments: []Comment{comment3, comment4}}
	users := []User{mary, ivan}
	filePath := "/home/george/goproject/src/student.vkusvill.ru/George/go-course/golang_study_ak/course1/13.popular_package/9.package_yaml/task1.13.9.3/test/file.yaml"
	err := writeYAML(filePath, users)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

type User struct {
	Name     string    `yaml:"name"`
	Age      int       `yaml:"age"`
	Comments []Comment `yaml:"comments"`
}
type Comment struct {
	Text string `yaml:"text"`
}

func writeYAML(filePath string, data []User) error {
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
