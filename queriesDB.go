package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

type jsonGemeinde struct {
	Gemeindename string `json:"gemeindename"`
}

var dbConn = connectDB()

func getGemeinde(c *gin.Context) {
	var response []jsonGemeinde
	var id = c.Query("id")

	stmt, err := dbConn.Prepare("SELECT gemeindename FROM public.gemeinde WHERE gkz=$1 LIMIT 1")
	checkError(err)
	rows, err := stmt.Query(id)
	checkError(err)

	for rows.Next() {
		var gemeinde string

		err = rows.Scan(&gemeinde)
		checkError(err)

		response = append(response, jsonGemeinde{Gemeindename: gemeinde})
	}

	c.IndentedJSON(http.StatusOK, response)
}
