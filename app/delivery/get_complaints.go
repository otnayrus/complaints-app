package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (h *handler) GetComplaints(c *gin.Context) {
	var (
		err error

		ctx = c.Request.Context()
	)

	complaints, err := h.repo.GetComplaints(ctx)
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	c.JSON(http.StatusOK, model.GetComplaintsResponse{
		Complaints: complaints,
	})

}
