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

type jsonStrasse struct {
	SKZ          int    `json:"skz"`
	Strassenname string `json:"strassenname"`
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
	var err error
	var dbConn *sql.DB
	// open database
	dbConn, err = sql.Open("postgres", psqlconn)
	checkError(err)

	// check db
	err = dbConn.Ping()
	checkError(err)

	return dbConn
}

// GET: http://localhost:7000/
func home(c *gin.Context) {
	var response = jsonMessage{
		Message: "Hello World! I'm the Go API.",
	}

	c.IndentedJSON(http.StatusOK, response)
}

// POST: http://localhost:7000/createTable
// BODY:
func createTable(c *gin.Context) {
	var sqlQuery = `DROP TABLE IF EXISTS public.strasse;

CREATE TABLE IF NOT EXISTS public.strasse
(
    skz integer NOT NULL,
    strassenname character varying(100) COLLATE pg_catalog."default",
    CONSTRAINT strasse_pkey PRIMARY KEY (skz)
)`
	_, err := dbConn.Query(sqlQuery)
	checkError(err)

	c.Status(http.StatusCreated)
}

// POST: http://localhost:7000/addStrasse
// HEADER: Content-Type: application/json
// BODY: {"skz":108711,"strassenname":"Andromedastraße"}
func addStrasse(c *gin.Context) {
	var strasse jsonStrasse

	err := c.BindJSON(&strasse)
	checkError(err)

	var sqlQuery = "INSERT INTO public.strasse(skz, strassenname) VALUES ($1, $2);"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	_, err = stmt.Query(strasse.SKZ, strasse.Strassenname)
	checkError(err)

	c.Status(http.StatusCreated)
}

// PUT: http://localhost:7000/changeStrasse/108711
// HEADER: Content-Type: application/json
// BODY: {"strassenname":"Andromedastraße2"}
func changeStrasse(c *gin.Context) {
	skz, err := strconv.Atoi(c.Param("skz"))
	checkError(err)

	var strasse jsonStrasse

	err = c.BindJSON(&strasse)
	checkError(err)

	var sqlQuery = "UPDATE public.strasse SET strassenname=$1 WHERE skz=$2;"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	_, err = stmt.Query(strasse.Strassenname, skz)
	checkError(err)

	c.Status(http.StatusOK)
}

// GET: http://localhost:7000/getStrasse?skz=108711
func getStrasse(c *gin.Context) {
	skz, err := strconv.Atoi(c.Query("skz"))
	checkError(err)

	var sqlQuery = "SELECT strassenname FROM public.strasse WHERE skz=$1 LIMIT 1;"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	rows, err := stmt.Query(skz)
	checkError(err)

	rows.Next()
	var strassenname string
	err = rows.Scan(&strassenname)

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, jsonStrasse{SKZ: skz, Strassenname: strassenname})
}

// DELETE: http://localhost:7000/deleteStrasse/108711
func deleteStrasse(c *gin.Context) {
	skz, err := strconv.Atoi(c.Param("skz"))
	checkError(err)

	var sqlQuery = "DELETE FROM public.strasse WHERE skz=$1;"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	_, err = stmt.Query(skz)
	checkError(err)

	c.Status(http.StatusOK)
}
