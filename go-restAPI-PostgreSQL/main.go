package main

import (
	"github.com/Michle99/goRestApiPostgreSQL/controllers"
	"github.com/Michle99/goRestApiPostgreSQL/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	db := models.SetupModels()

	// Provide db variables to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook) // update by id
	r.DELETE("/books/:id", controllers.DeleteBook) // delete by id
	r.Run()
}