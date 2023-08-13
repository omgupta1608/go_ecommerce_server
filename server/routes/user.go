package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/handlers"
)

func initUserRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/all", handlers.AllUsersHandler)
	}
}