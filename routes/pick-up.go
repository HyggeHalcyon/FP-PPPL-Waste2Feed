package routes

import (
	"Waste2Feed/config"
	"Waste2Feed/constants"
	"Waste2Feed/controller"
	"Waste2Feed/middleware"

	"github.com/gin-gonic/gin"
)

func PickUp(route *gin.Engine, pickupController controller.UserController, jwtService config.JWTService) {

	routes := route.Group("/api/pick-up")
	{
		routes.POST("", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_USER))
		routes.GET("", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_USER))
	}

	route.GET("/pick-up", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_USER))
}
