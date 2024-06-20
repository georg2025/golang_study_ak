package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "/home/george/goproject/src/student.vkusvill.ru/George/go-course/golang_study_ak/course1/13.popular_package/7.package_os/task1.13.7.3/test/file.txt"
	fmt.Println(ReadString(filePath))
}

// Our function
func ReadString(filePath string) string {
	// Here we got information from file
	content, err := os.ReadFile(filePath)
	// If we got error - we print it and return blank string
	if err != nil {
		fmt.Println("We have error with file reading:", err)
		return ""
	}
	// We return info from the file
	return string(content)
}
