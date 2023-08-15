package main

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/db"
	"github.com/omgupta1608/aftershoot_task/middlewares"
	"github.com/omgupta1608/aftershoot_task/routes"
	"github.com/omgupta1608/aftershoot_task/utils"
	_ "github.com/omgupta1608/aftershoot_task/cfg"
)

func main() {
	// initialize router
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	// connect to db
	if err := db.Connect(); err != nil {
		utils.PrintToConsole("Error in connecting to database: " + err.Error(), "error")
		return
	}

	// attach cors middlware
	server.Use(middlewares.CORSMiddleware())

	router := server.Group("/api/" + utils.GetVersion() + "/")

	// initialize routes
	routes.InitPublicRoutes(router)
	routes.InitPrivateRoutes(router)

	// schedule cron jobs
	utils.ScheduleJobs()

	server.Run(":5000")
}
