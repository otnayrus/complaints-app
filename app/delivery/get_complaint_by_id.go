package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (h *handler) GetComplaintByID(c *gin.Context) {
	var (
		err error
		req model.GetComplaintByIDRequest

		ctx = c.Request.Context()
	)

	err = c.BindUri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	complaint, err := h.repo.GetComplaintByID(ctx, req.ID)
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	c.JSON(http.StatusOK, model.GetComplaintByIDResponse{
		Complaint: *complaint,
	})
}
