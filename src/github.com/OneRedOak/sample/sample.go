package sample

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

func GetResult() string {

	// Establish db connection
	db, err := sql.Open("postgres", "user=postgres dbname=test sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	// db, err := sql.Open("postgres","host=/var/run/postgresql dbname=test user=postgres sslmode=disable")
	// defer db.Close()

	// Create table & column
	result, err := db.Exec("CREATE TABLE users (name varchar(255));")
	if err != nil {
		fmt.Println(err)
	}

	// Insert data into db
	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec("Daniel")
	if err != nil {
		fmt.Println(err, result, res)
	}

	// Retreive data from db
	var name string
	err = db.QueryRow("select name from users where name = ?", "Daniel").Scan(&name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(name)
	return "Daniel"
}