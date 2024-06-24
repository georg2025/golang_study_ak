package main

import (
	"reflect"
	"testing"
)

type tc struct {
	inData   []byte
	expected []User
	wantErr  bool
}

func Test_getJSON(t *testing.T) {
	comment1 := Comment{Text: "hello"}
	comment2 := Comment{Text: "go"}
	mary := User{Name: "Mary", Age: 32, Comments: []Comment{comment1, comment2}}
	users1 := []User{mary}
	users2 := []User{}
	testCases := []tc{
		{expected: users1, inData: []byte(`[{"name":"Mary","age":32,"comments":[{"text":"hello"},{"text":"go"}]}]`), wantErr: false},
		{expected: users2, inData: []byte{}, wantErr: false},
	}

	for _, tc := range testCases {
		result, err := getUsersFromJSON(tc.inData)

		if (err != nil) != tc.wantErr {
			t.Errorf("Got %v, expected %v", err, tc.wantErr)
		}

		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Got %v, expected %v", result, tc.expected)
		}
	}
}
