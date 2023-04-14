package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	const (
		host = "127.0.0.1"
		//host     = "192.168.0.2" // container ip
		port     = 5432
		user     = "postgres"
		password = "xsmmsgbAMfIOIWPPBrsc"
		database = "ogd"
	)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	var err error
	var dbConn *sql.DB
	// open database
	dbConn, err = sql.Open("postgres", psqlconn)
	checkError(err)

	// close database
	//defer db.Close()

	// check db
	err = dbConn.Ping()
	checkError(err)

	return dbConn
}
