package main

import (
	"reflect"
	"testing"
)

type tc struct {
	inData   []byte
	expected Person
	wantErr  bool
}

func TestUnmarshal(t *testing.T) {
	jsonData := []byte(`{"name":"John","age":30}`)
	yamlData := []byte(`name: Mary
age: 24`)
	jsonPerson := Person{Name: "John", Age: 30}
	yamlPerson := Person{Name: "Mary", Age: 24}
	testCases := []tc{
		{inData: jsonData, expected: jsonPerson, wantErr: false},
		{inData: yamlData, expected: yamlPerson, wantErr: false},
	}

	for _, tc := range testCases {
		result := Person{}
		err := unmarshal(tc.inData, &result)
		if (err != nil) != tc.wantErr {
			t.Errorf("Got %v, expected %v", err, tc.wantErr)
		}

		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Got %v, expected %v", result, tc.expected)
		}
	}
}
