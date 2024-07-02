package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/mattn/go-sqlite3"
)

// Определение структуры пользователя
type User struct {
	tableName string
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}
type Tabler interface {
	TableName() string
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenarateFakeUser() User
}

type GoFakeitGenerator struct {
	currentId int
}

type SQLiteGenerator struct {
}

func (user *User) TableName() string {
	return user.tableName
}

func (generator *SQLiteGenerator) CreateTableSQL(user *User) string {
	name := user.TableName()
	if name == "" {
		name = "test"
	}

	exec := fmt.Sprintf(`CREATE TABLE %s (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		email VARCHAR(100) UNIQUE
	)`, name)

	return exec
}

func (generator *SQLiteGenerator) CreateInsertSQL(user *User) string {
	exec := fmt.Sprintf("INSERT INTO %s (id, first_name, last_name, email) VALUES (?, ?, ?, ?)", user.TableName())

	return exec
}

func (generator *GoFakeitGenerator) GenerateFakeUser() User {
	user := User{}
	generator.currentId++
	user.ID = generator.currentId
	user.FirstName = gofakeit.FirstName()
	user.LastName = gofakeit.LastName()
	user.Email = gofakeit.Email()
	return user
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}
	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)
	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
}
