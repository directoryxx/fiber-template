package main

import (
	_ "clean-arch-template/docs"
	"clean-arch-template/internal/delivery/http"
	"clean-arch-template/internal/repository"
	"clean-arch-template/internal/usecase"
	"clean-arch-template/pkg/config"
	"clean-arch-template/pkg/database/sqlc"
	"clean-arch-template/pkg/infrastructure"
	"clean-arch-template/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"os"
)

// @title Clean Arch template
// @version 1.0
// @description API Docs Clean Arch Template
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1
func main() {
	newLogger := logger.NewLogger()

	/**
	Load Config
	*/
	envSource := config.LoadConfig()
	newLogger.Info("Loaded Config : " + envSource)

	/**
	Load Infrastructure
	*/
	newLogger.Info("Opening Connection Database")
	dbPool, err := infrastructure.OpenDB()
	if err != nil {
		newLogger.Error("Error Connection Database : " + err.Error())
		panic(err)
	}
	newLogger.Info("Success Connection Database")
	defer dbPool.Close()

	/**
	SQLC Init
	*/
	sqlcQueries := sqlc.New(dbPool)

	/**
	Init framework
	*/
	app := fiber.New(fiber.Config{
		ServerHeader: "App",
		AppName:      "Clean Arch Template v0.0.1",
	})

	/**
	Pprof
	*/
	//app.Use(pprof.New())

	// Create repository instances
	roleRepository := repository.NewRoleRepository(sqlcQueries)

	// Create usecase instances
	roleUsecase := usecase.NewRoleService(roleRepository)

	// Create handler instances
	handlers := &http.Handlers{
		//UserHandler: http.NewUserHandler(),
		RoleHandler: http.NewRoleHandler(roleUsecase),
		// Initialize other handlers as neededd
	}

	// Create router instance and pass the handler
	router := http.NewRouter(app, handlers)

	// Set up routes
	router.SetupRoutes()

	// Init Application
	err = app.Listen(":" + os.Getenv("APP_PORT"))

	if err != nil {
		newLogger.Error(err.Error())
	}
}
