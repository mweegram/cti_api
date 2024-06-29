package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Indicator struct {
	ID      int
	Type    string
	Value   string
	Comment string
	Date    string
	Actor   int
}

type Actor struct {
	ID      int
	Name    string
	Aliases []string
}

func Database_Connect() (*pgx.Conn, error) {
	_ = godotenv.Load(".env")
	db, err := pgx.Connect(context.Background(), os.Getenv("DB_STRING"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
