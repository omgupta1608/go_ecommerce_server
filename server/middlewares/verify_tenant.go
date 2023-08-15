package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/utils"
)

func VerifyTenants(tenants []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.Find(tenants, c.GetString("tenant_type")) == -1 {
			utils.SendError(c, http.StatusUnauthorized, errors.New("You are not authorized to use this API"))
			c.Abort()
		}
		c.Next()
	}
}
