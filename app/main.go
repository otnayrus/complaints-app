package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/delivery"
	"github.com/otnayrus/sb-rest/app/repository"
)

func main() {
	dbDsn := os.Getenv("DATABASE_URL")
	repo := repository.New(dbDsn)

	handler := delivery.New(repo)

	r := gin.Default()
	r.POST("/users", handler.CreateUser)
	r.PATCH("/users", handler.UpdateUser)
	r.DELETE("/users", handler.DeleteUser)

	r.POST("/categories", handler.CreateCategory)
	r.GET("/categories", handler.GetCategories)
	r.GET("/categories/:id", handler.GetCategoryByID)
	r.PATCH("/categories", handler.UpdateCategory)
	r.DELETE("/categories", handler.DeleteCategory)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.Run(":8001")
}
