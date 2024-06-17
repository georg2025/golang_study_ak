package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	data := `server: 
 port: "60"
db:
 host: 127.0.0.1
 port: "60"
 user: Ivanov
 password: Petrov`
	config, err := getConfigFromYAML([]byte(data))
	if err != nil {
		fmt.Println("we got error:", err)
	}
	fmt.Println(config)
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

func getConfigFromYAML(data []byte) (Config, error) {
	config := Config{}
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
