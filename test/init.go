package test

import (
	"clean-arch-template/pkg/config"
	"clean-arch-template/pkg/infrastructure"
	"clean-arch-template/pkg/logger"
	"context"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"strings"
	"time"
)

var newLogger = logger.NewLogger()

func initInfrastructure() {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:12",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "passwordq",
			"POSTGRES_DB":       "template",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5 * time.Second),
	}
	postgresContainer, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	endpointPostgres, _ := postgresContainer.Endpoint(ctx, "")

	dbInfo := strings.Split(endpointPostgres, ":")
	os.Setenv("DB_HOST", dbInfo[0])
	os.Setenv("DB_PORT", dbInfo[1])
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "passwordq")
	os.Setenv("DB_NAME", "template")
}

func migrationUp() {
	/**
	Load Config
	*/
	config.LoadConfig()

	dsn := infrastructure.GenerateDSN("postgres://")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		newLogger.Error(err.Error())
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		newLogger.Error(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../pkg/database/migrations",
		"postgres", driver)
	if err != nil {
		newLogger.Error(err.Error())
	}
	err = m.Up()
	if err != nil {
		return
	} // or m.Step(2) if you want to explicitly set the number of migrations to run
}
