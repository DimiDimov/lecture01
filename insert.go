package main

import (
	"database/sql"
	// "fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	row := os.Args[1:]

	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Post VALUES (NULL, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4])
	if err != nil {
		log.Fatal(err)
	}

}
