package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (h *handler) UpdateCategory(c *gin.Context) {
	var (
		err error
		req model.UpdateCategoryRequest

		ctx = c.Request.Context()
	)

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing, err := h.repo.GetCategoryByID(ctx, req.ID)
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	err = h.repo.UpdateCategory(ctx, req.MakeModel(*existing))
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      req.ID,
		"message": "Category updated successfully",
	})
}
