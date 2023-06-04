package http

import (
	"clean-arch-template/internal/repository"
	"clean-arch-template/internal/usecase"
	"clean-arch-template/pkg/config"
	"clean-arch-template/pkg/database/sqlc"
	"clean-arch-template/pkg/infrastructure"
	"clean-arch-template/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

var newLogger = logger.NewLogger()

func Run() *fiber.App {
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
	//defer dbPool.Close()

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
	handlers := &Handlers{
		//UserHandler: http.NewUserHandler(),
		RoleHandler: NewRoleHandler(roleUsecase),
		// Initialize other handlers as neededd
	}

	// Create router instance and pass the handler
	router := NewRouter(app, handlers)

	// Set up routes
	router.SetupRoutes()

	return app
}
