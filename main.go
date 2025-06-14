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

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func main() {
	var (
		db         *gorm.DB          = config.SetUpDatabaseConnection()
		jwtService config.JWTService = config.NewJWTService()

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		roleRepository repository.RoleRepository = repository.NewRoleRepository(db)

		userService service.UserService = service.NewUserService(userRepository, roleRepository)

		userController controller.UserController = controller.NewUserController(userService, jwtService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	server.LoadHTMLGlob("templates/*")

	routes.User(server, userController, jwtService)

	server.Static("/static", "./static")

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
