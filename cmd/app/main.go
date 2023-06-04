package main

import (
	_ "clean-arch-template/docs"
	"clean-arch-template/internal/delivery/http"
	"clean-arch-template/pkg/logger"
	"os"
)

var newLogger = logger.NewLogger()

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
	app := http.Run()

	// Init Application
	err := app.Listen(":" + os.Getenv("APP_PORT"))

	if err != nil {
		newLogger.Error(err.Error())
	}
}
