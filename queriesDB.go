package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

type jsonGemeinde struct {
	Gemeindename string `json:"gemeindename"`
}

var db *sql.DB

func initDB() {
	const (
		//host = "127.0.0.1"
		host     = "192.168.0.2" // container ip
		port     = 5432
		user     = "postgres"
		password = "xsmmsgbAMfIOIWPPBrsc"
		database = "ogd"
	)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	var err error
	// open database
	db, err = sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	//defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)
}

func getGemeinde(c *gin.Context) {
	if db == nil {
		initDB()
	}

	var response []jsonGemeinde
	id := c.Query("id")
	var sqlQuery = "SELECT gemeindename FROM public.gemeinde WHERE gkz=" + id + " LIMIT 1"

	rows, err := db.Query(sqlQuery)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string

		err = rows.Scan(&name)
		CheckError(err)

		response = append(response, jsonGemeinde{Gemeindename: name})
	}

	CheckError(err)
	c.IndentedJSON(http.StatusOK, response)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
