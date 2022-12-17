package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func doRegister(user string, password string) error {
	// Open a connection to the database
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the insert statement
	stmt, err := db.Prepare("INSERT INTO users (user, password) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the insert statement
	_, err = stmt.Exec(user, password)
	if err != nil {
		return err
	}

	return nil
}