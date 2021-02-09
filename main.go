// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"wxmp/routers/routers.go"
)


func main() {
	r := gin.Default()
	r.GET("/", routers.VerifySinagure)
	r.Run() // listen and serve on 0.0.0.0:8080
}