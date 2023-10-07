package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

// CreateCategory is a handler for creating a new category
func (h *handler) CreateCategory(c *gin.Context) {
	var (
		err error
		req model.CreateCategoryRequest

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

	id, err := h.repo.CreateCategory(ctx, req.MakeModel())
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
