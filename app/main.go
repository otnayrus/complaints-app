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
	r.POST("/users/login", handler.Login)
	r.POST("/users", handler.CreateUser)
	r.PATCH("/users", handler.IsAuthorizedUser(), handler.UpdateUser)
	r.DELETE("/users", handler.IsAuthorizedUser(), handler.DeleteUser)

	r.POST("/categories", handler.IsAuthorizedUser(), handler.IsAuthorizedAdmin(), handler.CreateCategory)
	r.GET("/categories", handler.GetCategories)
	r.GET("/categories/:id", handler.GetCategoryByID)
	r.PATCH("/categories", handler.IsAuthorizedUser(), handler.IsAuthorizedAdmin(), handler.UpdateCategory)
	r.DELETE("/categories", handler.IsAuthorizedUser(), handler.IsAuthorizedAdmin(), handler.DeleteCategory)

	r.POST("/complaints", handler.IsAuthorizedUser(), handler.CreateComplaint)
	r.GET("/complaints", handler.GetComplaints)
	r.GET("/complaints/:id", handler.GetComplaintByID)
	r.PATCH("/complaints", handler.IsAuthorizedUser(), handler.UpdateComplaint)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.Run(":8001")
}
