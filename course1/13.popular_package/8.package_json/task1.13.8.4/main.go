package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
	data := []map[string]interface{}{
		{"name": "Elliot",
			"age": 25},
	}
	filePath := "/home/george/goproject/src/student.vkusvill.ru/George/go-course/golang_study_ak/course1/13.popular_package/8.package_json/task1.13.8.4/test/file.txt"
	err := writeJSON(filePath, data)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func writeJSON(filePath string, data interface{}) error {
	path := filepath.Dir(filePath)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	infoToWrite, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, []byte(infoToWrite), 0644)
	return err
}
