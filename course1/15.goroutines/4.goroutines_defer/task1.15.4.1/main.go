package main

import (
	"os"
)

func main() {
}

func writeToFile(file *os.File, data string) error {
	defer file.Close()
	_, err := file.Write([]byte(data))

	if err != nil {
		return err
	}

	return nil
}
