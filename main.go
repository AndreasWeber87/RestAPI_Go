// Sources:
// https://go.dev/doc/tutorial/web-service-gin
// https://golangdocs.com/golang-postgresql-example

package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host = "127.0.0.1"
	// host = "192.168.0.2" // container ip
	port     = 5432
	user     = "postgres"
	password = "xsmmsgbAMfIOIWPPBrsc"
	database = "ogd"
)

var db *sql.DB

func main() {
	fmt.Println("Server started on port 9000...")
	fmt.Println("")
	fmt.Println("Possible calls:")
	fmt.Println("http://localhost:9000/")
	fmt.Println("GET: http://localhost:9000/hello/ic20b050")

	//handleRequests()

	router := gin.Default()
	router.GET("/", home)
	router.GET("/hello/:name", sayHelloGet)
	router.POST("/hello", sayHelloPost)

	router.Run("localhost:9000")

	//selectFromDB()
}

func initDB() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)
}

func selectFromDB() {
	if db == nil {
		initDB()
	}

	rows, err := db.Query(`SELECT gemeindename FROM public.gemeinde WHERE gkz=10101 LIMIT 1`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string

		err = rows.Scan(&name)
		CheckError(err)

		fmt.Println(name)
	}

	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
