package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/handlers"
)

func initAuthRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("/auth") 
	{
		authRouter.POST("/login", handlers.LoginHandler)
		authRouter.POST("/register", handlers.RegisterHandler)
	}
}