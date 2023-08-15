package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/db"
	"github.com/omgupta1608/aftershoot_task/utils"
)

func GetTop3CustomersHander(c *gin.Context) {
	users, err := db.Conn.GetTop3Customers(c)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Our top 3 customers", map[string]any{
		"customers": users,
	})

}
