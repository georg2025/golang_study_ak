package main

import (
	"reflect"
	"testing"
)

type tc struct {
	inData   []byte
	expected Config
	wantErr  bool
}

func TestGetConfigFromYAML(t *testing.T) {
	data := `server: 
 port: "60"
db:
 host: 127.0.0.1
 port: "60"
 user: Ivanov
 password: Petrov`
	server := Server{Port: "60"}
	db := Db{Host: "127.0.0.1", Port: "60", User: "Ivanov", Password: "Petrov"}
	config := Config{Server: server, Db: db}
	testCases := []tc{
		{expected: config, inData: []byte(data), wantErr: false},
	}

	for _, tc := range testCases {
		result, err := getConfigFromYAML(tc.inData)
		if (err != nil) != tc.wantErr {
			t.Errorf("Got %v, expected %v", err, tc.wantErr)
		}

		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Got %v, expected %v", result, tc.expected)
		}
	}
}
