package main

// EXECUTE "go get github.com/go-sql-driver/mysql" before run this go file
// EXECUTE "docker-compose up" as well

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Row struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySQL create and insert")

	// Open up our database connection.
	// A database is configured on docker
	// The database is called go_db
	db, err := sql.Open("mysql", "go_user:1234@tcp(127.0.0.1:3396)/go_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	createTableTest(db)

	insertRow(db)
	insertRow(db)
	insertRow(db)

	selectData(db)
}

func createTableTest(db *sql.DB) {
	// perform a db.Query create table
	result, err := db.Query("create table if not exists test (id int AUTO_INCREMENT, a_varchar varchar(255), PRIMARY KEY(ID))")
	// if there is an error creating, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer result.Close()
}

func insertRow(db *sql.DB) {
	result, err := db.Query("INSERT INTO test (a_varchar) VALUES ( 'TEST' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer result.Close()
}

func selectData(db *sql.DB) {
	// Execute the query
	results, err := db.Query("SELECT id, a_varchar FROM test")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var row Row
		// for each row, scan the result into our row struct
		err = results.Scan(&row.ID, &row.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the row's Name attribute
		log.Printf("%d - %s", row.ID, row.Name)
	}
}
