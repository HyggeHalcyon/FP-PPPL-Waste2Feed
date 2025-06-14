package routes

import (
	"Waste2Feed/config"
	"Waste2Feed/constants"
	"Waste2Feed/controller"
	"Waste2Feed/middleware"

	"github.com/gin-gonic/gin"
)

func Fodder(route *gin.Engine, userController controller.UserController, jwtService config.JWTService) {

	routes := route.Group("/api/fodder")
	{
		routes.POST("", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER))
		routes.GET("", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER))
	}

	route.GET("/fodder", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER))
}
