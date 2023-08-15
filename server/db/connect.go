package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Conn *Queries

func Connect() error {
	godotenv.Load()
	db, err := sql.Open("postgres", fmt.Sprintf("host=db user=postgres dbname=postgres sslmode=disable password=%s", os.Getenv("DATABASE_PASSWORD")))
	if err != nil {
		return err
	}

	Conn = New(db)
	return nil
}
