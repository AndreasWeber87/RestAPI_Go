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
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("Server started on port 7000...")
	fmt.Println("")
	fmt.Println("Possible calls:")
	fmt.Println("GET: http://localhost:7000/")
	fmt.Println("")
	fmt.Println("POST: http://localhost:7000/createTable")
	fmt.Println("	BODY:")
	fmt.Println("")
	fmt.Println("POST: http://localhost:7000/addStrasse")
	fmt.Println("	HEADER: Content-Type: application/json")
	fmt.Println("	BODY: {\"skz\":108711,\"strassenname\":\"Andromedastraße\"}")
	fmt.Println("")
	fmt.Println("PUT: http://localhost:7000/changeStrasse/108711")
	fmt.Println("	HEADER: Content-Type: application/json")
	fmt.Println("	BODY: {\"strassenname\":\"Andromedastraße2\"}")
	fmt.Println("")
	fmt.Println("GET: http://localhost:7000/getStrasse?skz=108711")
	fmt.Println("")
	fmt.Println("DELETE: http://localhost:7000/deleteStrasse/108711")

	router := gin.Default()
	router.GET("/", home)

	router.POST("/createTable", createTable)
	router.POST("/addStrasse", addStrasse)
	router.PUT("/changeStrasse/:skz", changeStrasse)
	router.GET("/getStrasse", getStrasse)
	router.DELETE("/deleteStrasse/:skz", deleteStrasse)

	err := router.Run(":7000")
	//err := router.Run(":10000")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
