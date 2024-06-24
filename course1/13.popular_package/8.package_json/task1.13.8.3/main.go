package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Text string `json:"text"`
}

func main() {
	comment1 := Comment{Text: "hello"}
	comment2 := Comment{Text: "go"}
	comment3 := Comment{Text: "moo"}
	comment4 := Comment{Text: "what"}
	mary := User{Name: "Mary", Age: 32, Comments: []Comment{comment1, comment2}}
	ivan := User{Name: "Ivan", Age: 21, Comments: []Comment{comment3, comment4}}
	users := []User{mary, ivan}
	filePath := "/home/george/goproject/src/student.vkusvill.ru/George/go-course/golang_study_ak/course1/13.popular_package/8.package_json/task1.13.8.3/test/file.txt"
	err := WriteFile(filePath, users)

	if err != nil {
		fmt.Println(err)
	}
}

func getJSON(data []User) (string, error) {
	if len(data) == 0 {
		return "", nil
	}

	var sb strings.Builder

	for _, i := range data {
		jsonData, err := json.Marshal(i)

		if err != nil {
			return "", err
		}

		sb.WriteString(string(jsonData))
	}

	return sb.String(), nil
}

func WriteFile(filePath string, data []User) error {
	path := filepath.Dir(filePath)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}

	infoToWrite, err := getJSON(data)

	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, []byte(infoToWrite), 0644)
	return err
}
