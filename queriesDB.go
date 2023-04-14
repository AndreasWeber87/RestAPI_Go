package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

type jsonGemeinde struct {
	GKZ          int    `json:"gkz"`
	Gemeindename string `json:"gemeindename"`
}

var dbConn = connectDB()

func createTable(c *gin.Context) {
	var sqlQuery = `DROP TABLE IF EXISTS public.gemeinde;

CREATE TABLE IF NOT EXISTS public.gemeinde
(
    gkz integer NOT NULL,
    gemeindename character varying(50) COLLATE pg_catalog."default",
    CONSTRAINT gemeinde_pkey PRIMARY KEY (gkz)
)`
	_, err := dbConn.Query(sqlQuery)
	checkError(err)

	c.Status(http.StatusCreated)
}

func addGemeinde(c *gin.Context) {
	var gemeinde jsonGemeinde

	err := c.BindJSON(&gemeinde)
	checkError(err)

	var sqlQuery = "INSERT INTO public.gemeinde(gkz, gemeindename) VALUES ($1, $2);"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	_, err = stmt.Query(gemeinde.GKZ, gemeinde.Gemeindename)
	checkError(err)

	c.Status(http.StatusCreated)
}

func changeGemeinde(c *gin.Context) {
	var gemeinde jsonGemeinde

	err := c.BindJSON(&gemeinde)
	checkError(err)

	var sqlQuery = "UPDATE public.gemeinde SET gemeindename=$1 WHERE gkz=$2;"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	_, err = stmt.Query(gemeinde.Gemeindename, gemeinde.GKZ)
	checkError(err)

	c.Status(http.StatusOK)
}

func getGemeinde(c *gin.Context) {
	gkz, err := strconv.Atoi(c.Query("gkz"))
	checkError(err)

	var sqlQuery = "SELECT gemeindename FROM public.gemeinde WHERE gkz=$1 LIMIT 1"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	rows, err := stmt.Query(gkz)
	checkError(err)

	rows.Next()
	var gemeindename string
	err = rows.Scan(&gemeindename)

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, jsonGemeinde{GKZ: gkz, Gemeindename: gemeindename})
}

func deleteGemeinde(c *gin.Context) {
	gkz, err := strconv.Atoi(c.Param("gkz"))
	checkError(err)

	var sqlQuery = "DELETE FROM public.gemeinde WHERE gkz=$1;"
	stmt, err := dbConn.Prepare(sqlQuery)
	checkError(err)

	_, err = stmt.Query(gkz)
	checkError(err)

	c.Status(http.StatusOK)
}
