package main

import (
	"reflect"
	"testing"
)

func TestFilterWords(t *testing.T) {
	censorMap := map[string]string{
		"foo": "moo",
		"poo": "doo",
	}
	result := filterWords("Foo koo zoo! Poo foo hoo!", censorMap)
	expected := "Moo koo zoo! Doo moo hoo!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	result = filterWords("", censorMap)
	expected = ""
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSplitSentences(t *testing.T) {
	result := splitSentences("foo!")
	expected := []string{"foo"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, GOT %v", expected, result)
	}
	result = splitSentences("foo moo! moo!")
	expected = []string{"foo moo", " moo"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, GOT %v", expected, result)
	}
}

func TestCheckUpper(t *testing.T) {
	result := checkUpper("Foo", "moo")
	expected := "Moo"
	if result != expected {
		t.Errorf("Expected %s, GOT %s", expected, result)
	}
	result = checkUpper("", "")
	expected = ""
	if result != expected {
		t.Errorf("Expected %s, GOT %s", expected, result)
	}
	result = checkUpper("foo", "moo")
	expected = "moo"
	if result != expected {
		t.Errorf("Expected %s, GOT %s", expected, result)
	}
}

func TestWordsToSentence(t *testing.T) {
	result := WordsToSentence([]string{"moo", "foo", "boo"})
	expected := "moo foo boo!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
