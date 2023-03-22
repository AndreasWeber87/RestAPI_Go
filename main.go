// Sources:
// https://go.dev/doc/tutorial/web-service-gin
// https://golangdocs.com/golang-postgresql-example
// https://dev.to/umschaudhary/blog-project-with-go-gin-mysql-and-docker-part-1-3cg1

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
	fmt.Println("GET: http://localhost:7000/hello/ic20b050")
	fmt.Println("POST: http://localhost:7000/hello  name=ic20b050")
	fmt.Println("")
	fmt.Println("GET: http://localhost:7000/getGemeinde/10101")

	router := gin.Default()
	router.GET("/", home)

	router.GET("/hello/:name", sayHelloGet)
	router.POST("/hello", sayHelloPost)

	router.GET("/getGemeinde/:id", getGemeinde)

	router.Run(":7000")
}
