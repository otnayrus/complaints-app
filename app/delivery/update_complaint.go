package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (h *handler) UpdateComplaint(c *gin.Context) {
	var (
		err error
		req model.UpdateComplaintRequest

		ctx = c.Request.Context()
	)

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = req.Validate(h.validator)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing, err := h.repo.GetComplaintByID(ctx, req.ID)
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	userID := c.GetInt("user_id")
	rolesRaw, exists := c.Get("roles")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "roles not found"})
		return
	}
	rolesMap, ok := rolesRaw.(map[string]bool)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "roles conversion failed"})
		return
	}

	complaint, err := req.MakeModel(*existing, userID, rolesMap[string(model.RoleAdmin)])
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	err = h.repo.UpdateComplaint(ctx, complaint)
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      req.ID,
		"message": "Complaint updated successfully",
	})
}
