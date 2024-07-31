package main

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

func main() {
}
