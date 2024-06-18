package main

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
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

	_, err = db.Exec(`CREATE TABLE users (
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

	request, args, err := PrepareQuery("insert", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(request, args...)
	if err != nil {
		return err
	}

	return nil
}

func SelectUser(user User) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	request, args, err := PrepareQuery("select", "users", user)
	if err != nil {
		return User{}, err
	}
	row := db.QueryRow(request, args...)
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

	request, args, err := PrepareQuery("update", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(request, args...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	request, args, err := PrepareQuery("delete", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(request, args...)
	if err != nil {
		return err
	}

	return nil
}

func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
	sql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)
	var request string
	var args []interface{}
	var err error
	switch operation {
	case "insert":
		request, args, err = sql.Insert(table).Columns("name", "age").Values(user.Name, user.Age).ToSql()
	case "update":
		request, args, err = sql.Update(table).Set("age", user.Age).Where(squirrel.Eq{"name": user.Name}).ToSql()
	case "select":
		request, args, err = sql.Select("id", "name", "age").From(table).Where(squirrel.Eq{"id": user.ID}).ToSql()
	case "delete":
		request, args, err = sql.Delete(table).Where(squirrel.Eq{"id": user.ID}).ToSql()
	default:
		return "", nil, fmt.Errorf("cant figure out what operation do you want")
	}
	return request, args, err
}
