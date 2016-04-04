
# Info340b &mdash; Lecture 01
## Overview
The examples included in this repo introduce the basic concepts of reading and writing to a SQLite3 database before combining these concepts into a  simple _screen scraper_ that grabs posts from the Vertabelo bog and loads into our database.

## Setup
Prior to running the examples, you will need to run the included script to create tables. You can do this from the commandline `sqlite3 blog.db < create.sql` or from within the SQLite3 shell using `.open` to load/create the database file and `.read` to execute the commands in the .sql file.

A similarly named `drop.sql` is also included to remove the tables from the database, though you might find it just as easy to delete the file itself and recreate.

## Running Examples
Go examples can easily be run from the command line or through your IDE. From the command line, we use `go run select.go` to build and run the program (in this case our select.go example).

Since Go is a compiled language, it\'s also likely that we\'ll want to build in one step so that we can run the program natively. If a program excepts command line arguments, this is going to be the best strategy for testing (and you may find it easier to do from the terminal). Here\'s an example using insert.go, which inserts a row into the database based on our commandline arguments (careful ... we don\'t have any error checking):
```
> go build insert.go
> ./insert author title 2016-04-03 http://test.com "summary text"
```
The _./_ syntax is specific to *nix. Go on Windows compiles the program into the _.exe_ format natively such that you can just run `insert` followed by the proper arguments.