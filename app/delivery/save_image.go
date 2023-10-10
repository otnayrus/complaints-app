package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handler) SaveImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": "File upload failed"})
		return
	}

	path := uuid.New().String() + file.Filename
	err = c.SaveUploadedFile(file, "uploads/"+path)
	if err != nil {
		c.JSON(500, gin.H{"error": "File save failed"})
		return
	}

	c.JSON(200, gin.H{
		"message": "File upload success",
		"path":    "http://localhost:8000/images/" + path,
	})
}
