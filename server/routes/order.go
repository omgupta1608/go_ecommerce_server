package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/handlers"
)

func initOrderRoutes(router *gin.RouterGroup) {
	orderRouter := router.Group("/order")
	{
		orderRouter.POST("/new", handlers.PlaceOrderHandler)
		orderRouter.POST("/process", handlers.ProcessOrderHandler)
	}
}