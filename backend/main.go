package main

import (
	"TicketManagement/config"
	"TicketManagement/middleware"
	"TicketManagement/router"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	config.LoadDB()

	config.ConnectRedis()

	r := gin.Default()

	rl := middleware.NewRateLimiter(50, time.Minute)
	r.Use(rl.Middleware())
	r.Use(middleware.CORSMiddleware())
	r.Static("/public", "./public")

	api := r.Group("/api")

	router.AuthRouter(api)
	router.DepartmentRouter(api)
	router.RoleRouter(api)
	router.TicketTypeRouter(api)
	router.UserRouter(api)
	router.WorkflowPathRouter(api)
	router.WorkflowRouter(api)
	router.TicketRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.Port))
}
