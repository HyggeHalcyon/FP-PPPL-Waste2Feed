package routes

import (
	"Waste2Feed/config"
	"Waste2Feed/constants"
	"Waste2Feed/controller"
	"Waste2Feed/middleware"

	"github.com/gin-gonic/gin"
)

func Fodder(route *gin.Engine, fodderController controller.FodderController, jwtService config.JWTService) {

	routes := route.Group("/api/fodder")
	{
		routes.POST("", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER), fodderController.Create)
		routes.POST("/pay", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER), fodderController.Pay)
	}

	routes = route.Group("/fodder")
	{
		routes.GET("", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER), fodderController.ViewFodder)
		routes.GET("/payment", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER), fodderController.ViewPaymentPage)
		routes.GET("/pending", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER), fodderController.ViewPendingPayments)
		routes.GET("/order", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER), fodderController.ViewOrder)
		routes.GET("/information", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER), fodderController.ViewFodderInformation)
		routes.GET("/nutrition", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER), fodderController.ViewNutrition)
		routes.GET("/dosage", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_FARMER), fodderController.ViewDosage)
	}
}
