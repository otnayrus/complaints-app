package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/delivery"
	"github.com/otnayrus/sb-rest/app/repository"
)

func main() {
	dbDsn := os.Getenv("DATABASE_URL")
	userRepo := repository.New(dbDsn)

	handler := delivery.New(userRepo)

	r := gin.Default()
	r.POST("/users", handler.CreateUser)
	r.PATCH("/users", handler.UpdateUser)
	r.DELETE("/users", handler.DeleteUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
