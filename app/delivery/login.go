package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
	"github.com/otnayrus/sb-rest/app/pkg/jwt"
	"github.com/otnayrus/sb-rest/app/pkg/secret"
)

func (h *handler) Login(c *gin.Context) {
	var (
		req model.LoginRequest
		err error

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

	user, err := h.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	err = secret.MatchPassword(req.Password, user.Password)
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	roles, err := h.repo.GetUserRoles(ctx, user.ID)
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	var role string = string(model.RoleUser)
	if roles[string(model.RoleAdmin)] {
		role = string(model.RoleAdmin)
	}

	token, err := jwt.GenerateJWTStringWithClaims(map[string]interface{}{
		"user_id": user.ID,
		"name":    user.Name,
		"role":    role,
	})
	if err != nil {
		c.JSON(errorwrapper.ConvertToHTTPError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
