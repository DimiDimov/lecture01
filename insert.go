package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	/*
	 The OS package grants us access to a structure containing the arguments passed from the command line. This
	 structure is known as a "slice" in Golang and is closely related to the array. The odd looking index passed is
	 shorthand to create a new slice ranging from the second element (slices are 0-indexed) to the end. In an Args
	 slice, the first element is the command itself. We're after the actual arguments that follow.

	 Note the := syntax we're using here. Go requires variables to be declared with a datatype before we can use them.
	 The language does provide a shortcut that we can use when we are ready to assign the value at the time of
	 of variable declaration. In these scenarios, Go can infer the type.

	 // longform (Go needs to know the type)
	 var class string
	 class = "info340b"

	 // shortform (Go infers that type is string from context)
	 class := "info340b"
	 */
	row := os.Args[1:]

	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a prepared statement for our INSERT into the database
	// id (autoincrement), author, title, date, link, summary
	stmt, err := db.Prepare("INSERT INTO Post VALUES (NULL, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	// Execute prepared statement
	_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4])
	if err != nil {
		log.Fatal(err)
	}

}
