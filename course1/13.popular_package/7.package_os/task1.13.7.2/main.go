package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filePath := "/home/george/goproject/src/student.vkusvill.ru/George/go-course/golang_study_ak/course1/13.popular_package/7.package_os/task1.13.7.2/test/file.txt"
	path := filepath.Dir(filePath)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	err = WriteFile(strings.NewReader("Hello, World!"), file)
	if err != nil {
		fmt.Println("error with writing on file:", err)
	}
}

func WriteFile(data io.Reader, fd io.Writer) error {
	_, err := io.Copy(fd, data)

	return err
}
