// Sources:
// https://go.dev/doc/tutorial/web-service-gin
// https://golangdocs.com/golang-postgresql-example
// https://dev.to/umschaudhary/blog-project-with-go-gin-mysql-and-docker-part-1-3cg1
// https://go.dev/doc/database/prepared-statements

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server started on port 7000...")
	fmt.Println("")
	fmt.Println("Possible calls:")
	fmt.Println("GET: http://localhost:7000/")
	fmt.Println("")
	fmt.Println("POST: http://localhost:7000/createTable")
	fmt.Println("	BODY:")
	fmt.Println("")
	fmt.Println("POST: http://localhost:7000/addStreet")
	fmt.Println("	HEADER: Content-Type: application/json")
	fmt.Println("	BODY: {\"skz\":108711,\"streetname\":\"Andromedastraße\"}")
	fmt.Println("")
	fmt.Println("PUT: http://localhost:7000/changeStreet/108711")
	fmt.Println("	HEADER: Content-Type: application/json")
	fmt.Println("	BODY: {\"streetname\":\"Andromedastraße2\"}")
	fmt.Println("")
	fmt.Println("GET: http://localhost:7000/getStreet?skz=108711")
	fmt.Println("")
	fmt.Println("DELETE: http://localhost:7000/deleteStreet/108711")

	router := gin.Default()
	router.GET("/", home)

	router.POST("/createTable", createTable)
	router.POST("/addStreet", addStreet)
	router.PUT("/changeStreet/:skz", changeStreet)
	router.GET("/getStreet", getStreet)
	router.DELETE("/deleteStreet/:skz", deleteStreet)

	err := router.Run(":7000")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
