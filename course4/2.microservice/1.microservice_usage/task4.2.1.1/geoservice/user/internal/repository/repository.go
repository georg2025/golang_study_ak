package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	models "user/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgressDataBase struct {
	Base *sql.DB
}

type UserRepository interface {
	GetByID(ctx context.Context, id string) (models.User, error)
}

func StartPostgressDataBase(ctx context.Context) (*PostgressDataBase, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	dataBase := &PostgressDataBase{}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return dataBase, err
	}

	dataBase.Base = db
	err = dataBase.CreateNewUserTable()
	return dataBase, err
}

func (db *PostgressDataBase) CreateNewUserTable() error {
	newTableString := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL,
		isExist BOOLEAN DEFAULT true
	);`

	_, err := db.Base.Exec(newTableString)
	return err
}

func (db *PostgressDataBase) GetByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	query := `SELECT username, password FROM users WHERE id = $1`

	row := db.Base.QueryRow(query, id)
	err := row.Scan(&user.Username, &user.Password)

	if err != nil {

		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user with ID %s not found", id)
		}

		return user, err
	}

	return user, nil
}
