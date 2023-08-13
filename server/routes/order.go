package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/handlers"
)

func initOrderRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/order")
	{
		userRouter.POST("/new", handlers.PlaceOrderHandler)
	}
}