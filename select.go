package main

/*

 Imports can be conducted one per line with the import keyword on each, but most Golang programmers will import
 a list all at once:

 1. Open and close with parens
 2. Surround each package in double quotes on its own line
 3. Do not include any other internal punctuation like commas or semicolons
 4. Only reference packages that are currently in use
 5. If you don't need to reference a package by name, e.g., the SQLite driver, prefix it with the _ operator

 */

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

/*
 Every standalone go app has a main function, which the compiler uses to find the starting point of execution.
 */
func main() {
	/*
	 In almost every language or framework, you'll discover the need to open a database connection (or often a pool
	 of connections) before executing queries. This setup is typically done once to begin followed by a single
	 teardown at the end of execution. Connection setup can be slow, so it's generally a good idea to keep it open
	 until you're sure you won't need it again. In a more complex applilcation, we'd pass the sql.DB value to
	 functions that need to use it.

	 Most IO commands follow the multiple return pattern, including a response and an error. We should always check
	 the error value -- probably a bit more thoroughly than we do here. If you choose to ignore the value during
	 development, the compiler will require that you substitue the _ operator in place of a named error variable.
	 */
	db, err := sql.Open("sqlite3", "./blog.db")

	/*
	 We won't do anything here unless the call to Open returned a real value for err. Remember that nil is a special
	 place holder that means "no value".
	 */
	if err != nil {
		log.Fatal(err)
	}
	/*
	 Since our application is only a single function, we'll make deferred call to Close here. It will execute after
	 the function completes and all asynchronous calls return or tiemout.
	 */
	defer db.Close()

	/*
	 Same multi-return pattern with error, but this time we're going to get the result set.
	 */
	rows, err := db.Query("SELECT * FROM Post")
	if err != nil {
		log.Fatal(err)
	}

	/*
	 Due to the way query result sets operate, you need to free up resources once you're done with it. We actually
	 make the call to do this right away, though we're deferring execution until the function is complete and all
	 asynchronous calls have returned or timed out.
	 */
	defer rows.Close()

	/*
	 Just like we can import a list of packages, we can declare a list of variables all at once without retyping the
	 var keyword on every line, e.g.,

	 	var id int
	 	var author string
	 	...

	 Follow the same rules we listed above.
	 */
	var (
		id      int
		author  string
		title   string
		date    string
		link    string
		summary string
	)

	/*
	 No parens around the for loop conditions as you might expect. For is also the only loop contstruct in Golang. In
	 this situation, the loop executes until the rows object returns false (no more rows) or encounters an error.
	 */
	for rows.Next() {
		/*
		 The Scan function doesn't inherently have access to our variables, so we pass them "by reference" with the &
		 operator. In a nutshell, &id passes scan the memory location for the id variable. Scan can "dereference"
		 that variable with the * operator to write directly to the location. Technically this means that Scan is
		 creating a side-effect rather than actually returning a list of values back.

		 Don't worry if this is a bit confusing. Afterall, our focus is the database and SQL, not the nitty gritty of
		 Golang.

 		 If you're paying close attention, you might have noticed that Scan() must handle type conversions implicitly.
		 */
		rows.Scan(&id, &author, &title, &date, &link, &summary)

		/*
		 fmt is a pretty handy package that gives us access to the C-inspired Printf function. Provide a format
		 string and a list of variables to replace the placeholders (%d and %s in this example). You'll find plenty
		 of guidance online for using Printf effectively.
		 */
		fmt.Printf("%d | %s | %s\n", id, author, title)
	}

	/*
	 It's possible for errors to occur and kick us out of the loop above. We need to check this explicitly at the end
	 of the loop to ensure that we didn't exit prematurely on an error condition.
	 */
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
