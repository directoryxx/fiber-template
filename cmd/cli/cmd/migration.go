/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"clean-arch-template/pkg/config"
	"clean-arch-template/pkg/infrastructure"
	"clean-arch-template/pkg/logger"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "up" {
			migrationUp()
		} else if len(args) == 1 && args[0] == "down" {
			migrationDown()
		} else if len(args) == 2 && args[0] == "create" && len(args[1]) != 0 {
			createMigration(args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func migrationUp() {
	newLogger := logger.NewLogger()

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
		"file://pkg/database/migrations",
		"postgres", driver)
	if err != nil {
		newLogger.Error(err.Error())
	}
	err = m.Up()
	if err != nil {
		return
	} // or m.Step(2) if you want to explicitly set the number of migrations to run
}

func migrationDown() {
	newLogger := logger.NewLogger()

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
		"file://pkg/database/migrations",
		"postgres", driver)
	if err != nil {
		newLogger.Error(err.Error())
	}
	err = m.Down()
	if err != nil {
		return
	} // or m.Step(2) if you want to explicitly set the number of migrations to run
}

func createMigration(name string) {
	cmd := exec.Command("/go/bin/migrate", "create", "-ext", "sql", "-dir", "pkg/database/migrations", "-seq", name)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
