package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/omgupta1608/aftershoot_task/cfg"
	"github.com/omgupta1608/aftershoot_task/db"
	"github.com/omgupta1608/aftershoot_task/middlewares"
	"github.com/omgupta1608/aftershoot_task/routes"
	ws "github.com/omgupta1608/aftershoot_task/websocket"

	"github.com/omgupta1608/aftershoot_task/utils"
)

func main() {
	go ws.H.Run()
	
	// initialize router
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	// connect to db
	if err := db.Connect(); err != nil {
		utils.PrintToConsole("Error in connecting to database: "+err.Error(), "error")
		return
	}

	utils.PrintToConsole("Connected to Database!", "info")

	// attach cors middlware
	server.Use(middlewares.CORSMiddleware())

	// initialize websocket handler
	server.GET("/", func(ctx *gin.Context) {
		ws.ServeWs(ctx.Writer, ctx.Request)
	})

	router := server.Group("/api/" + utils.GetVersion() + "/")

	// initialize routes
	routes.InitPublicRoutes(router)
	routes.InitPrivateRoutes(router)

	// schedule cron jobs
	utils.ScheduleJobs()

	server.Run(":5000")
}
