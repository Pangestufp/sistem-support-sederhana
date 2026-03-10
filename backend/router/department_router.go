package router

import (
	"TicketManagement/config"
	"TicketManagement/handler"
	"TicketManagement/middleware"
	"TicketManagement/repository"
	"TicketManagement/service"

	"github.com/gin-gonic/gin"
)

func DepartmentRouter(api *gin.RouterGroup) {
	DepartmentRepository := repository.NewDepartmentRepository(config.DB)
	DepartmentService := service.NewDeparmentService(DepartmentRepository, config.RedisClient)
	DepartmentHandler := handler.NewDepartmentHandler(DepartmentService)

	department := api.Group("/department")

	department.Use(middleware.JWTMiddleware())

	department.POST("", middleware.RoleMiddleware(101), DepartmentHandler.Create)
	department.GET("", DepartmentHandler.GetAll)
	department.GET("/:id", DepartmentHandler.GetByID)
	department.PUT("/:id", middleware.RoleMiddleware(101), DepartmentHandler.Update)
	department.DELETE("/:id", middleware.RoleMiddleware(101), DepartmentHandler.Delete)
}
