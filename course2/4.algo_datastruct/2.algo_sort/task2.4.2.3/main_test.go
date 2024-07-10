package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	arr1 := []User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 3, Name: "Charlie", Age: 25},
	}
	arr2 := []User{
		{ID: 2, Name: "Bob", Age: 28},
		{ID: 4, Name: "David", Age: 35},
	}

	expected := []User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 28},
		{ID: 3, Name: "Charlie", Age: 25},
		{ID: 4, Name: "David", Age: 35},
	}

	result := Merge(arr1, arr2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMergeEmptyArrays(t *testing.T) {
	arr1 := []User{{ID: 1, Name: "Alice", Age: 30}}
	arr2 := []User{}

	expected := []User{{ID: 1, Name: "Alice", Age: 30}}

	result := Merge(arr1, arr2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	result = Merge(arr2, arr1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
