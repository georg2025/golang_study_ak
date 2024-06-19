package main

import (
	"reflect"
	"testing"
)

func TestGetUniqueUsers(t *testing.T) {
	user1 := User{Nickname: "Brad", Age: 22, Email: "loop@loop.ru"}
	user2 := User{Nickname: "Fred", Age: 28, Email: "loop1@loop.ru"}
	user3 := User{Nickname: "Brad", Age: 32, Email: "loop2@loop.ru"}
	users := []User{user1, user2, user3}
	result := getUniqueUsers(users)
	expected := []User{user1, user2}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Got %v, expected %v", result, expected)
	}
}
