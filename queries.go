package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

type jsonMessage struct {
	Message string `json:"message"`
}

type jsonStreet struct {
	SKZ        int    `json:"skz"`
	Streetname string `json:"streetname"`
}

var dbConn = connectDB()

func connectDB() *sql.DB {
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
	// open database
	dbConn, err := sql.Open("postgres", psqlconn)
	checkError(err)

	// check db
	err = dbConn.Ping()
	checkError(err)

	return dbConn
}

// GET: http://localhost:7000/
func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jsonMessage{Message: "Hello World! I'm the Go API."})
}

// POST: http://localhost:7000/createTable
// BODY:
func createTable(c *gin.Context) {
	const sqlQuery = `DROP TABLE IF EXISTS public.strasse;

CREATE TABLE IF NOT EXISTS public.strasse
(
    skz integer NOT NULL,
    strassenname character varying(100) COLLATE pg_catalog."default",
    CONSTRAINT strasse_pkey PRIMARY KEY (skz)
)`
	_, err := dbConn.Exec(sqlQuery)
	checkError(err)

	c.IndentedJSON(http.StatusCreated, jsonMessage{Message: "Table created successfully."})
}

// POST: http://localhost:7000/addStreet
// HEADER: Content-Type: application/json
// BODY: {"skz":108711,"streetname":"Andromedastraße"}
func addStreet(c *gin.Context) {
	var street jsonStreet

	err := c.BindJSON(&street)
	checkError(err)

	const sqlQuery = "INSERT INTO public.strasse(skz, strassenname) VALUES ($1, $2);"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	_, err = stmt.Exec(street.SKZ, street.Streetname)
	checkError(err)

	c.IndentedJSON(http.StatusCreated, jsonMessage{Message: "Street added successfully."})
}

// PUT: http://localhost:7000/changeStreet/108711
// HEADER: Content-Type: application/json
// BODY: {"streetname":"Andromedastraße2"}
func changeStreet(c *gin.Context) {
	skz, err := strconv.Atoi(c.Param("skz"))
	checkError(err)

	var street jsonStreet

	err = c.BindJSON(&street)
	checkError(err)

	const sqlQuery = "UPDATE public.strasse SET strassenname=$1 WHERE skz=$2;"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	_, err = stmt.Exec(street.Streetname, skz)
	checkError(err)

	c.IndentedJSON(http.StatusOK, jsonMessage{Message: "Street changed successfully."})
}

// GET: http://localhost:7000/getStreet?skz=108711
func getStreet(c *gin.Context) {
	skz, err := strconv.Atoi(c.Query("skz"))
	checkError(err)

	const sqlQuery = "SELECT strassenname FROM public.strasse WHERE skz=$1 LIMIT 1;"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	rows, err := stmt.Query(skz)
	checkError(err)

	rows.Next()
	var streetname string
	err = rows.Scan(&streetname)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, jsonMessage{Message: "No street found."})
		return
	}

	c.IndentedJSON(http.StatusOK, jsonStreet{SKZ: skz, Streetname: streetname})
}

// DELETE: http://localhost:7000/deleteStreet/108711
func deleteStreet(c *gin.Context) {
	skz, err := strconv.Atoi(c.Param("skz"))
	checkError(err)

	const sqlQuery = "DELETE FROM public.strasse WHERE skz=$1;"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	_, err = stmt.Exec(skz)
	checkError(err)

	c.IndentedJSON(http.StatusOK, jsonMessage{Message: "Street deleted successfully."})
}
