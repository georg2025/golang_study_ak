package main

import (
	"reflect"
	"testing"
)

func TestCommentsFunctions(t *testing.T) {
	users := []User{
		{
			Name: "Ivan",
			Comments: []Comment{
				{Message: "Good COmmEnT 1"},
				{Message: "BaD COmmEnT 1"},
				{Message: "Normal COmmEnT 1"},
				{Message: "Bad COmmEnT 2"},
			},
		},
		{
			Name: "Zoya",
			Comments: []Comment{
				{Message: "Bad COmmEnT 1"},
				{Message: "BaD COmmEnT 2"},
				{Message: "Normal COmmEnT 1"},
				{Message: "Bad COmmEnT 3"},
			},
		},
	}
	result := FilterComments(users)
	expected := []User{
		{
			Name: "Ivan",
			Comments: []Comment{
				{Message: "Good COmmEnT 1"},
				{Message: "Normal COmmEnT 1"},
			},
		},
		{
			Name: "Zoya",
			Comments: []Comment{
				{Message: "Normal COmmEnT 1"},
			},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Got %v, expected %v", result, expected)
	}
	badCommentsResult := GetBadComments(users)
	badCommentsExpected := []Comment{
		{Message: "BaD COmmEnT 1"},
		{Message: "Bad COmmEnT 2"},
		{Message: "Bad COmmEnT 1"},
		{Message: "BaD COmmEnT 2"},
		{Message: "Bad COmmEnT 3"},
	}
	if !reflect.DeepEqual(badCommentsResult, badCommentsExpected) {
		t.Errorf("Got %v, expected %v", badCommentsResult, badCommentsExpected)
	}
	goodCommentsResult := GetGoodComments(users)
	goodCommentsExpected := []Comment{
		{Message: "Good COmmEnT 1"},
	}
	if !reflect.DeepEqual(goodCommentsResult, goodCommentsExpected) {
		t.Errorf("Got %v, expected %v", goodCommentsResult, goodCommentsExpected)
	}

}
