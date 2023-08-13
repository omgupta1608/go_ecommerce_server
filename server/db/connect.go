package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Conn *Queries

func Connect() error {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=postgrespassword")
	if err != nil {
		return err
	}

	Conn = New(db)
	return nil
}
