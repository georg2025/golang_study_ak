package main

import (
	"fmt"
	"os/exec"
)

func main() {
	filePath := "/home/george/goproject/src/student.vkusvill.ru/George/go-course/golang_study_ak/course1/13.popular_package/7.package_os/task1.13.7.4/test/main"
	fmt.Println(ExecBin(filePath))
}

// Our function
func ExecBin(binPath string, args ...string) string {
	out, err := exec.Command(binPath, args...).Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(out)
}
