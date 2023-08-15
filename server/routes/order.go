package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/handlers"
	"github.com/omgupta1608/aftershoot_task/middlewares"
)

func initOrderRoutes(router *gin.RouterGroup) {
	orderRouter := router.Group("/order")
	{
		orderRouter.POST("/new", middlewares.VerifyTenants([]string{"CUSTOMER"}), handlers.PlaceOrderHandler)
		orderRouter.POST("/process", middlewares.VerifyTenants([]string{"ADMIN"}), handlers.ProcessOrderHandler)
		orderRouter.GET("/:id", handlers.GetOrderDetails)
	}
}
