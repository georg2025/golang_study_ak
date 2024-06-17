package main

import (
	"testing"
)

type tc struct {
	inData   []User
	expected string
	wantErr  bool
}

func Test_getJSON(t *testing.T) {
	comment1 := Comment{Text: "hello"}
	comment2 := Comment{Text: "go"}
	comment3 := Comment{Text: "moo"}
	comment4 := Comment{Text: "what"}
	mary := User{Name: "Mary", Age: 32, Comments: []Comment{comment1, comment2}}
	ivan := User{Name: "Ivan", Age: 21, Comments: []Comment{comment3, comment4}}
	users1 := []User{mary, ivan}
	users2 := []User{}
	testCases := []tc{
		{inData: users1, expected: `{"name":"Mary","age":32,"comments":[{"text":"hello"},{"text":"go"}]}{"name":"Ivan","age":21,"comments":[{"text":"moo"},{"text":"what"}]}`, wantErr: false},
		{inData: users2, expected: "", wantErr: false},
	}

	for _, tc := range testCases {
		result, err := getJSON(tc.inData)
		if (err != nil) != tc.wantErr {
			t.Errorf("Got %v, expected %v", err, tc.wantErr)
		}
		if result != tc.expected {
			t.Errorf("Got %s, expected %s", result, tc.expected)
		}
	}
}
