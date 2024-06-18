package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := CreateUserTable()
	if err != nil {
		fmt.Println("Error with User Table creating:", err)
	}
	user := User{ID: 1, Name: "Mary", Age: 22}
	InsertUser(user)
	fmt.Println(SelectUser(1))
	user2 := User{ID: 1, Name: "Mary", Age: 23}
	err = UpdateUser(user2)
	if err != nil {
		fmt.Println("Error with users updating:", err)
	}
	fmt.Println(SelectUser(1))
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE Users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	)`)

	if err != nil {
		return err
	}

	return nil
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
	if err != nil {
		return err
	}

	return nil
}

func SelectUser(id int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	request := fmt.Sprintf("SELECT * FROM users WHERE id = %d", id)
	row := db.QueryRow(request)
	answer := User{}

	err = row.Scan(&answer.ID, &answer.Name, &answer.Age)
	if err != nil {
		return answer, err
	}

	return answer, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users SET age = ? Where name = ?", user.Age, user.Name, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
