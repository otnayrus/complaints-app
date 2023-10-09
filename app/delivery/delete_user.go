package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (h *handler) DeleteUser(c *gin.Context) {
	var (
		req model.DeleteUserRequest
		err error

		ctx = c.Request.Context()
	)

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("user_id")
	if userID != req.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can't delete other users"})
		return
	}

	err = h.repo.DeleteUser(ctx, userID)
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      0,
		"message": "User deleted successfully",
	})

}
