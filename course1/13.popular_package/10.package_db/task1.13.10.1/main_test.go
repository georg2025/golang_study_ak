package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	err := CreateUserTable()
	if err != nil {
		t.Errorf("Error with User Table creating: %v", err)
	}

	user := User{ID: 1, Name: "Mary", Age: 22}
	InsertUser(user)
	result, err := SelectUser(1)
	if err != nil {
		t.Errorf("Error with selecting user: %v", err)
	}

	if !reflect.DeepEqual(user, result) {
		t.Errorf("Got %v, expected %v", result, user)
	}

	user2 := User{ID: 1, Name: "Mary", Age: 23}
	err = UpdateUser(user2)
	if err != nil {
		t.Errorf("Error with users updating: %v", err)
	}

	result, err = SelectUser(1)
	if err != nil {
		t.Errorf("Error with selecting user: %v", err)
	}

	if !reflect.DeepEqual(user2, result) {
		t.Errorf("Got %v, expected %v", result, user2)
	}

	err = DeleteUser(1)
	if err != nil {
		t.Errorf("Error with selecting user: %v", err)
	}

	_, err = SelectUser(1)
	if err == nil {
		t.Errorf("Expected error, got result")
	}
}
