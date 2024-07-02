package main

import (
	"testing"
)

func TestCreateTableSQLWithDefaultTableName(t *testing.T) {
	sqlGenerator := &SQLiteGenerator{}
	user := User{}
	expectedSQL := `CREATE TABLE test (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		email VARCHAR(100) UNIQUE
	)`
	actualSQL := sqlGenerator.CreateTableSQL(&user)

	if actualSQL != expectedSQL {
		t.Errorf("Got %s, expected %s", actualSQL, expectedSQL)
	}
}

func TestCreateTableSQLWithEmptyTableName(t *testing.T) {
	sqlGenerator := &SQLiteGenerator{}
	user := User{}
	expectedSQL := `CREATE TABLE test (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		email VARCHAR(100) UNIQUE
	)`
	actualSQL := sqlGenerator.CreateTableSQL(&user)

	if actualSQL != expectedSQL {
		t.Errorf("Got %s, expected %s", actualSQL, expectedSQL)
	}
}

func TestTabler(t *testing.T) {
	user := User{
		ID:        1,
		FirstName: "Egor",
		LastName:  "Ivanov",
		Email:     "ivanov@egor.ru",
	}

	result := user.FirstName
	expected := "Egor"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestCreateInsertSQL(t *testing.T) {
	user := User{
		ID:        1,
		FirstName: "Egor",
		LastName:  "Ivanov",
		Email:     "ivanov@egor.ru",
	}

	SqlGenerator := &SQLiteGenerator{}
	result := SqlGenerator.CreateInsertSQL(&user)
	expected := "INSERT INTO test (id, first_name, last_name, email) VALUES (?, ?, ?, ?)"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestGenerateFakeUserIncrementsID(t *testing.T) {
	generator := &GoFakeitGenerator{}

	user1 := generator.GenerateFakeUser()
	user2 := generator.GenerateFakeUser()

	if user1.ID != 1 || user2.ID != 2 {
		t.Errorf("User1 ID should be 1, user2 ID should be 2, but actually user1 IS = %d, user2 ID = %d",
			user1.ID, user2.ID)
	}
}
