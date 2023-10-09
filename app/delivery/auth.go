package delivery

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
	jwtUtil "github.com/otnayrus/sb-rest/app/pkg/jwt"
)

const ContextUserIDKey = "user_id"
const ContextRolesKey = "roles"

func (h *handler) IsAuthorizedUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, err := jwtUtil.ParseJWTStringWithClaims(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
				"error":   err.Error(),
			})
			return
		}

		userID := int(claim[ContextUserIDKey].(float64))
		c.Set(ContextUserIDKey, userID)
	}
}

func (h *handler) IsAuthorizedAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt(ContextUserIDKey)
		log.Println(userID, string(model.RoleAdmin), "IsAuthorizedAdmin")

		isAdmin, err := h.repo.IsUserHaveRole(c.Request.Context(), userID, string(model.RoleAdmin))
		if err != nil {
			c.JSON(errorwrapper.ConvertToHTTPError(err))
			return
		}

		if !isAdmin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}
	}
}

func (h *handler) GetAuthorizedRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt(ContextUserIDKey)
		roles, err := h.repo.GetUserRoles(c.Request.Context(), userID)
		if err != nil {
			c.AbortWithStatusJSON(errorwrapper.ConvertToHTTPError(err))
			return
		}

		c.Set(ContextRolesKey, roles)
	}
}
