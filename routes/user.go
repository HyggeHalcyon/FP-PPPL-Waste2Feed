package routes

import (
	"Waste2Feed/config"
	"Waste2Feed/constants"
	"Waste2Feed/controller"
	"Waste2Feed/middleware"

	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService config.JWTService) {

	routes := route.Group("/api/user")
	{
		routes.POST("/login", userController.Login)
		routes.POST("/register", userController.Register)
		routes.GET("/me", middleware.AuthenticateBearer(jwtService), userController.Me)
		routes.PATCH("/update", middleware.AuthenticateBearer(jwtService), userController.Update)
		routes.POST("/redeem", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_USER), userController.Redeem)

	}

	routes = route.Group("/api/courier")
	{
		routes.POST("/add", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN), userController.NewCourier)
		routes.PATCH("/", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN), userController.EditCourier)
		routes.DELETE("/", middleware.AuthenticateBearer(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN), userController.DeleteCourier)
	}

	route.GET("/", userController.ViewIndex)
	route.GET("/register", userController.ViewRegister)
	route.GET("/login", userController.ViewLogin)
	route.GET("/choose-role", userController.ViewChooseRole)
	route.GET("/dashboard", middleware.AuthenticateCookies(jwtService), userController.ViewDashboard)
	route.GET("/redeem", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_USER), userController.ViewRedeem)
	route.GET("/profile", middleware.AuthenticateCookies(jwtService), userController.ViewProfile)
	route.GET("/faq", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_USER, constants.ENUM_ROLE_FARMER), userController.ViewFAQ)
	route.GET("/courier", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN), userController.ViewCourier)
	route.GET("/add-courier", middleware.AuthenticateCookies(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN), userController.RenderAddCourierPage)
}
