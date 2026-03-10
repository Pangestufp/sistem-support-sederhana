package router

import (
	"TicketManagement/config"
	"TicketManagement/handler"
	"TicketManagement/middleware"
	"TicketManagement/repository"
	"TicketManagement/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup) {
	UserRepository := repository.NewUserRepository(config.DB)
	UserService := service.NewUserService(UserRepository)
	UserHandler := handler.NewUserHandler(UserService)

	user := api.Group("/user")
	user.Use(middleware.JWTMiddleware())

	user.GET("", UserHandler.GetAll)
	user.GET("/own", UserHandler.GetBySelf)
	user.GET("/:id", UserHandler.GetByID)
	user.PUT("/:id", middleware.RoleMiddleware(101), UserHandler.Update)
	user.PATCH("/:id/password", middleware.RoleMiddleware(101), UserHandler.UpdatePassword)
	user.DELETE("/:id", middleware.RoleMiddleware(101), UserHandler.Delete)

}
