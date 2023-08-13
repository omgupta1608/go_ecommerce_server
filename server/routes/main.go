package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/middlewares"
)

func InitPublicRoutes(router *gin.RouterGroup) {
	initAuthRoutes(router)
}

func InitPrivateRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.JWTMiddleware())
	initUserRoutes(router)
	initProductRoutes(router)
	initOrderRoutes(router)
}