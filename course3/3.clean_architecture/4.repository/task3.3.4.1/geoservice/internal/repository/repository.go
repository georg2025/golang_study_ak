package repository

import (
	"context"
	"fmt"
	service "geoservice/internal/service"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgressDataBase struct {
	Pool *pgxpool.Pool
}

type UserRepository interface {
	Create(ctx context.Context, user service.User) error
	GetByID(ctx context.Context, id string) (service.User, error)
	Update(ctx context.Context, user service.User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]service.User, error)
}

func StartPostgressDataBase(ctx context.Context) error {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	dataBase := NewPostgressDataBase(pool)
	newTableString := `CREATE TABLE IF NOT EXISTS newtable (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(32) NOT NULL,
		isExist BOOLEAN NOT NULL
	);`
	dataBase.Pool.Exec(context.Background(), newTableString)

	<-ctx.Done()
	fmt.Println("we stop gracefuly")
}

func NewPostgressDataBase(pool *pgxpool.Pool) *PostgressDataBase {
	dataBase := &PostgressDataBase{Pool: pool}
	return dataBase
}
