package main

import (
	"os"
	"path/filepath"
)

func main() {
	err := WriteFile("/home/george/goproject/src/student.vkusvill.ru/George/go-course/golang_study_ak/course1/13.popular_package/7.package_os/task1.13.7.1/test/file.txt",
		[]byte("Hello, World!"), os.FileMode(0644))
	if err != nil {
		panic("We are panicing")
	}
}

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	path := filepath.Dir(filePath)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	err = os.WriteFile(filePath, data, perm)
	return err
}
