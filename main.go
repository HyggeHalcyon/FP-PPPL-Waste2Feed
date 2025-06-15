package main

import (
	"log"
	"os"

	"Waste2Feed/config"
	"Waste2Feed/controller"
	"Waste2Feed/middleware"
	"Waste2Feed/migrations/seeder"
	"Waste2Feed/repository"
	"Waste2Feed/routes"
	"Waste2Feed/service"
	"Waste2Feed/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func main() {
	var (
		db         *gorm.DB          = config.SetUpDatabaseConnection()
		jwtService config.JWTService = config.NewJWTService()

		userRepository   repository.UserRepository   = repository.NewUserRepository(db)
		roleRepository   repository.RoleRepository   = repository.NewRoleRepository(db)
		pickupRepository repository.PickupRepository = repository.NewPickupRepository(db)
		fodderRepository repository.FodderRepository = repository.NewFodderRepository(db)

		userService   service.UserService   = service.NewUserService(userRepository, roleRepository)
		pickupService service.PickupService = service.NewPickupService(pickupRepository)
		fodderService service.FodderService = service.NewFodderService(fodderRepository)

		userController   controller.UserController   = controller.NewUserController(userService, jwtService)
		pickupController controller.PickupController = controller.NewPickupController(pickupService, jwtService)
		fodderController controller.FodderController = controller.NewFodderController(fodderService, jwtService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	server.Static("/static", "./static")

	utils.SetUpFuncs(server)
	server.LoadHTMLGlob("templates/*")

	routes.User(server, userController, jwtService)
	routes.Pickup(server, pickupController, jwtService)
	routes.Fodder(server, fodderController, jwtService)

	if err := seeder.RunSeeders(db); err != nil {
		log.Fatalf("error migration seeder: %v", err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
