package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/mattn/go-sqlite3"
)

// Определение структуры пользователя
type User struct {
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

type DataWriter interface {
	WriteData(db *sql.DB, data ...Tabler) error
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenarateFakeUser() User
}

type Migrate interface {
	Migrate(data ...Tabler) error
}

type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLGenerator
}

func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	return &Migrator{
		db:           db,
		sqlGenerator: sqlGenerator,
	}
}

type GoFakeitGenerator struct {
	currentId int
}

type SQLiteGenerator struct {
}

func (user *User) TableName() string {
	return "test"
}

func (generator *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	name := table.TableName()

	exec := fmt.Sprintf(`CREATE TABLE %s (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		email VARCHAR(100) UNIQUE
	)`, name)

	return exec
}

func (generator *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	exec := fmt.Sprintf("INSERT INTO %s (id, first_name, last_name, email) VALUES (?, ?, ?, ?)",
		model.TableName())

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

func (generator *SQLiteGenerator) WriteData(db *sql.DB, data ...Tabler) error {
	if len(data) == 0 {
		return nil
	}
	_, err := db.Exec(generator.CreateTableSQL(data[0].(*User)))
	if err != nil {
		fmt.Println("error with table creation: ", err)
	}

	for _, user := range data {
		_, err = db.Exec(generator.CreateInsertSQL(user.(*User)), user.(*User).ID,
			user.(*User).FirstName, user.(*User).LastName, user.(*User).Email)

		fmt.Println(generator.CreateInsertSQL(user.(*User)))

		if err != nil {
			return err
		}
	}

	return nil
}

func (generator *SQLiteGenerator) ReadData(db *sql.DB, tableName string) ([]Tabler, error) {
	users := []Tabler{}
	selectData := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(selectData)

	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

		if err != nil {
			return users, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (migrator *Migrator) Migrate(db *sql.DB, data ...Tabler) error {
	for _, user := range data {
		users, err := migrator.sqlGenerator.(*SQLiteGenerator).ReadData(db, user.TableName())

		if err != nil {
			fmt.Println("error with taking data: ", err)
		}

		err = migrator.sqlGenerator.(*SQLiteGenerator).WriteData(migrator.db, users...)

		if err != nil {
			fmt.Println("error with writing data: ", err)
		}

	}

	return nil
}

func main() {
	sqlGenerator := &SQLiteGenerator{}

	oldDb, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("Cant open database")
	}
	defer oldDb.Close()

	newDb, err := sql.Open("sqlite3", "users.db")

	if err != nil {
		fmt.Println("Cant open database")
	}
	defer newDb.Close()

	migrator := NewMigrator(newDb, sqlGenerator)

	if err := migrator.Migrate(oldDb, &User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
