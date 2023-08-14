package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/handlers"
)

func initProductRoutes(router *gin.RouterGroup) {
	productRouter := router.Group("/product")
	{
		productRouter.POST("/new", handlers.AddNewProductHandler)
		productRouter.GET("/", handlers.GetProductsHandler)
		productRouter.POST("/rate", handlers.RateProductHandler)
	}
}
