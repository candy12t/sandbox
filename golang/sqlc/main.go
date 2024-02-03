package main

import (
	"database/sql"
	"github/candy12t/sandbox/golang/sqlc/tutorial"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := sql.Open("mysql", "")
	if err != nil {
		return err
	}

	queries := tutorial.New(db)
	return nil
}
