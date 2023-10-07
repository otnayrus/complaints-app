package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

// GetCategories is a handler for getting all categories
func (h *handler) GetCategories(c *gin.Context) {
	var (
		err error

		ctx = c.Request.Context()
	)

	categories, err := h.repo.GetCategories(ctx)
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	c.JSON(http.StatusOK, model.GetCategoriesResponse{
		Categories: categories,
	})
}
