package main

import (
	"database/sql"
	"log"

	_ "github.com/apache/calcite-avatica-go"
)

func main() {
	// Connections are defined by a DSN
	// The format is http://address:port[/schema][?parameter1=value&...parameterN=value]
	// For COD, BASIC authentication is used.
	// The workload username and password are passed as parameters avaticaUser and avaticaPassword
	//
	// For example:
	// COD URL: 'https://gateway.env.cloudera.site/cdp-proxy-api/avatica/'
	// Workload username: jgoodson
	// Workload password: Secret1!
	// Would result in this DSN:
	dsn := "https://gateway.env.cloudera.site/cdp-proxy-api/avatica/?&authentication=BASIC&avaticaUser=jgoodson&avaticaPassword=Secret1!"

	log.Println("Connecting...")
	db, err := sql.Open("avatica", dsn)
	if err != nil {
		log.Fatal("Connection: ", err)
	}
	defer db.Close()

	log.Println("Create table if not exists...")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username VARCHAR)")
	if err != nil {
		log.Fatal("Create: ", err)
	}

	log.Println("Insert a row...")
	_, err = db.Exec("UPSERT INTO users VALUES (?, ?)", 1, "admin")
	if err != nil {
		log.Println("Insert: ", err)
	}

	log.Println("Reading and printing rows...")
	var (
		id       int
		username string
	)
	rows, err := db.Query("SELECT id, username from users")
	if err != nil {
		log.Fatal("Query: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &username)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, username)
	}
}
