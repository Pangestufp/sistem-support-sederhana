package router

import (
	"TicketManagement/config"
	"TicketManagement/handler"
	"TicketManagement/repository"
	"TicketManagement/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	AuthRepository := repository.NewAuthRepository(config.DB)
	UserRoleRepository := repository.NewUserRoleRepository(config.DB)
	AuthService := service.NewAuthService(AuthRepository, UserRoleRepository)
	AuthHandler := handler.NewAuthHandler(AuthService)

	api.POST("/register", AuthHandler.Register)
	api.POST("/login", AuthHandler.Login)
}
