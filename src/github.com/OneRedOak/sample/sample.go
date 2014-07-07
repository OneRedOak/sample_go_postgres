package sample

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
)

func GetResult() string {

	// Establish db connection
	db, err := sql.Open("postgres", "user=postgres dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Create table & column
	result, err := db.Exec("CREATE TABLE users (name varchar(255));")
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into db
	stmt, err := db.Prepare("INSERT INTO users(name) VALUES($1)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Daniel")
	if err != nil {
		log.Fatal(err, result, res)
	}

	// Retreive data from db
	var name string
	err = db.QueryRow("select name from users where name = $1", "Daniel").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}