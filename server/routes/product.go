package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/handlers"
)

func initProductRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/product")
	{
		userRouter.POST("/new", handlers.AddNewProductHandler)
	}
}