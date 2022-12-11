package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	handleRequests()
}
