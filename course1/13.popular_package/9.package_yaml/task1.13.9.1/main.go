package main

import (
	"bytes"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	server := Server{Port: "80"}
	db := Db{Host: "127.0.0.1", Port: "80", User: "Petrovich", Password: "Sidorovich"}
	config := Config{Server: server, Db: db}
	server2 := Server{Port: "60"}
	db2 := Db{Host: "127.0.0.1", Port: "60", User: "Ivanov", Password: "Petrov"}
	config2 := Config{Server: server2, Db: db2}
	info, err := getYAML([]Config{config, config2})

	if err != nil {
		fmt.Println("we have error:", err)
	}
	fmt.Println(info)
}

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func getYAML(config []Config) (string, error) {
	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	err := encoder.Encode(config)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
