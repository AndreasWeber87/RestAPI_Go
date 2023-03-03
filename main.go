// Sources:
// https://go.dev/doc/tutorial/web-service-gin
// https://golangdocs.com/golang-postgresql-example

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server started on port 9000...")
	fmt.Println("")
	fmt.Println("Possible calls:")
	fmt.Println("http://localhost:9000/")
	fmt.Println("GET: http://localhost:9000/hello/ic20b050")
	fmt.Println("POST: http://localhost:9000/hello  name=ic20b050")
	fmt.Println("")
	fmt.Println("GET: http://localhost:9000/getGemeinde/10101")

	router := gin.Default()
	router.GET("/", home)

	router.GET("/hello/:name", sayHelloGet)
	router.POST("/hello", sayHelloPost)

	router.GET("/getGemeinde/:id", getGemeinde)

	router.Run("localhost:9000")
}
