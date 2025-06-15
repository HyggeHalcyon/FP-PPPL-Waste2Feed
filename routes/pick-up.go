package routes

import (
	"Waste2Feed/config"
	"Waste2Feed/constants"
	"Waste2Feed/controller"
	"Waste2Feed/middleware"

	"github.com/gin-gonic/gin"
)

func Pickup(route *gin.Engine, pickupController controller.PickupController, jwtService config.JWTService) {

	routes := route.Group("/api/pick-up")
	{
		routes.POST("", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_USER), pickupController.Create)
	}

	route.GET("/pick-up", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_USER), pickupController.ViewPickup)
	route.GET("/history", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_USER, constants.ENUM_ROLE_COURIER), pickupController.ViewHistory)
}
