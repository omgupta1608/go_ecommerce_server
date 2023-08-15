package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/handlers"
	"github.com/omgupta1608/aftershoot_task/middlewares"
)

func initProductRoutes(router *gin.RouterGroup) {
	productRouter := router.Group("/product")
	{
		productRouter.POST("/new", middlewares.VerifyTenants([]string{"ADMIN"}), handlers.AddNewProductHandler)
		productRouter.GET("/", handlers.GetProductsHandler)
		productRouter.POST("/rate", middlewares.VerifyTenants([]string{"CUSTOMER"}), handlers.RateProductHandler)
	}
}
