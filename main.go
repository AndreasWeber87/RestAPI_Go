// Sources:
// https://go.dev/doc/tutorial/web-service-gin
// https://golangdocs.com/golang-postgresql-example

package main

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"
)

const (
	host = "127.0.0.1"
	// host = "192.168.0.2" // container ip
	port     = 5432
	user     = "postgres"
	password = "xsmmsgbAMfIOIWPPBrsc"
	database = "ogd"
)

func returnHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"message\":\"Hello World! I'm the Go API.\"}")
	//json.NewEncoder(w).Encode(Articles)
}

func returnName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := strings.TrimPrefix(r.URL.RawQuery, "/hello")
	fmt.Fprintf(w, "{\"message\":\"Hello "+strings.Split(name, "=")[1]+"! I'm the Go API.\"}")
}

func handleRequests() {
	http.HandleFunc("/", returnHello)
	http.HandleFunc("/hello", returnName)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("Server started on port 9000...")
	fmt.Println("")
	fmt.Println("Possible calls:")
	fmt.Println("http://localhost:9000/")
	fmt.Println("GET: http://localhost:9000/hello?name=ic20b050")

	//handleRequests()

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

	fmt.Println("Connected!")
	selectFromDB(db)
}

func selectFromDB(db *sql.DB) {
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
