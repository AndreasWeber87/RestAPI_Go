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
	fmt.Println("http://localhost:7000/")
	fmt.Println("POST: http://localhost:7000/createTable")
	fmt.Println("POST: http://localhost:7000/addGemeinde   gkz=10101 gemeindename=Eisenstadt")
	fmt.Println("PUT: http://localhost:7000/changeGemeinde   gkz=10101 gemeindename=Eisenstadt2")
	fmt.Println("GET: http://localhost:7000/getGemeinde?gkz=10101")
	fmt.Println("DELETE: http://localhost:7000/deleteGemeinde/10101")

	router := gin.Default()
	router.GET("/", home)

	router.POST("/createTable", createTable)
	router.POST("/addGemeinde", addGemeinde)
	router.PUT("/changeGemeinde", changeGemeinde)
	router.GET("/getGemeinde", getGemeinde)
	router.DELETE("/deleteGemeinde/:gkz", deleteGemeinde)

	err := router.Run(":10000")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
