package main

import (
	"net/http"

	"github.com/Michle99/mysql_docker_container/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		infrastructure.LoadEnv()
		infrastructure.NewDatabase()
		context.JSON(http.StatusOK, gin.H{ "data": "Hello world!" })
	})

	router.Run(":4000")
}