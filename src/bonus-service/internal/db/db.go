package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"postgres", 5432, "postgres", "privileges", "postgres")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
