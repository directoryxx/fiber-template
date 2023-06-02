/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"clean-arch-template/pkg/helper"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// migrationCmd represents the migration command
var generatorCmd = &cobra.Command{
	Use:   "generatorcrud",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		if len(args) != 0 {
			fmt.Println("Generate Repository")
			generateRepository(args[0])
			fmt.Println("Generate Usecase")
			generateUsecase(args[0])
			fmt.Println("Generate Handler")
			generateHandler(args[0])
		} else {
			fmt.Println("Missing Parameter. Ex : go run cmd/cli/main.go generatorcrud acl")
		}
	},
}

func init() {
	rootCmd.AddCommand(generatorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateRepository(feat string) bool {
	// 1 capital
	// 2 normal
	sprintf := fmt.Sprintf(`package repository

import (
	"clean-arch-template/pkg/database/sqlc"
	"context"
	"strconv"
)

type %[1]sRepository interface {
	GetByID%[1]s(ctx context.Context, id string) (*sqlc.%[1]s, error)
	Create%[1]s(ctx context.Context, name string) (*sqlc.%[1]s, error)
	ListAll%[1]s(ctx context.Context) (*[]sqlc.%[1]s, error)
	Delete%[1]s(ctx context.Context, id string) error
	Count%[1]s(ctx context.Context, id string) (count int32, err error)
	CountAll%[1]s(ctx context.Context) (count int32, err error)
	PaginationList%[1]s(ctx context.Context, limit int, offset int) (*[]sqlc.%[1]s, error)
	// other repository methods...
}

type %[2]sRepository struct {
	db *sqlc.Queries
}

func New%[1]sRepository(db *sqlc.Queries) %[1]sRepository {
	return &%[2]sRepository{
		db: db,
	}
}

func (r *%[2]sRepository) GetByID%[1]s(ctx context.Context, id string) (*sqlc.%[1]s, error) {
	idConv, _ := strconv.ParseInt(id, 10, 0)
	%[2]s, err := r.db.Get%[1]s(ctx, int32(idConv))
	if err != nil {
		return nil, err
	}
	return &%[2]s, nil
}

func (r *%[2]sRepository) ListAll%[1]s(ctx context.Context) (*[]sqlc.%[1]s, error) {
	%[2]sModel, err := r.db.List%[1]ss(ctx)
	if err != nil {
		return nil, err
	}
	return &%[2]sModel, nil
}

func (r *%[2]sRepository) Create%[1]s(ctx context.Context, name string) (*sqlc.%[1]s, error) {
	%[2]s, err := r.db.Create%[1]s(ctx, name)
	if err != nil {
		return nil, err
	}
	return &%[2]s, nil
}

func (r *%[2]sRepository) Delete%[1]s(ctx context.Context, id string) (err error) {
	idConv, _ := strconv.ParseInt(id, 10, 0)
	error := r.db.Delete%[1]s(ctx, int32(idConv))
	if error != nil {
		return error
	}

	return nil
}

func (r *%[2]sRepository) Count%[1]s(ctx context.Context, id string) (count int32, err error) {
	idConv, _ := strconv.ParseInt(id, 10, 0)
	count, err = r.db.Count%[1]sByID(ctx, int32(idConv))
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *%[2]sRepository) CountAll%[1]s(ctx context.Context) (count int32, err error) {
	count, err = r.db.Count%[1]sAll(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *%[2]sRepository) PaginationList%[1]s(ctx context.Context, limit int, offset int) (*[]sqlc.%[1]s, error) {
	args := sqlc.List%[1]ssPaginationParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	list%[1]s, err := r.db.List%[1]ssPagination(ctx, args)
	if err != nil {
		return nil, err
	}

	return &list%[1]s, err
}

		`, helper.CapitalizeWord(feat), feat)

	f, err := os.Create("./internal/repository/" + feat + ".repository.go")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(sprintf)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Created Repository to : ./internal/repository/" + feat + ".repository.go")
	return true
}

func generateUsecase(feat string) {
	sprintf := fmt.Sprintf(`package usecase

import (
	"clean-arch-template/internal/repository"
	"clean-arch-template/pkg/database/sqlc"
	"clean-arch-template/pkg/helper"
	"context"
)

type %[1]sUsecase interface {
	Get%[1]sByID(ctx context.Context, id string) (*sqlc.%[1]s, error)
	Create%[1]s(ctx context.Context, name string) (*sqlc.%[1]s, error)
	List%[1]s(ctx context.Context) (*[]sqlc.%[1]s, error)
	List%[1]sPagination(ctx context.Context, page int, page_size int) (*[]sqlc.%[1]s, int32, error)
	Delete%[1]s(ctx context.Context, id string) error
	// other use case methods...
}

type %[2]sUsecase struct {
	%[2]sRepo repository.%[1]sRepository
}

func New%[1]sService(%[2]sRepo repository.%[1]sRepository) %[1]sUsecase {
	return &%[2]sUsecase{%[2]sRepo: %[2]sRepo}
}

func (s *%[2]sUsecase) Get%[1]sByID(ctx context.Context, id string) (*sqlc.%[1]s, error) {
	u, err := s.%[2]sRepo.GetByID%[1]s(ctx, id)
	count, errCount := s.%[2]sRepo.Count%[1]s(ctx, id)
	if errCount != nil {
		return nil, err
	}

	if count == 0 {
		%[2]sEmpty := &sqlc.%[1]s{
			ID: 0,
		}
		return %[2]sEmpty, nil
	}

	if err != nil {
		return nil, err
	}

	// Perform any additional business logic or transformations
	return u, nil
}

func (s *%[2]sUsecase) List%[1]s(ctx context.Context) (*[]sqlc.%[1]s, error) {
	u, err := s.%[2]sRepo.ListAll%[1]s(ctx)
	if err != nil {
		return nil, err
	}
	// Perform any additional business logic or transformations
	return u, nil
}

func (s *%[2]sUsecase) List%[1]sPagination(ctx context.Context, page int, page_size int) (%[2]s *[]sqlc.%[1]s, count int32, err error) {
	offsetCalculate := helper.OffsetCalculator(page, page_size)
	limitCalculate := helper.LimitCalculator(page, page_size)
	u, err := s.%[2]sRepo.PaginationList%[1]s(ctx, limitCalculate, offsetCalculate)
	count%[1]s, errCount := s.%[2]sRepo.CountAll%[1]s(ctx)
	if errCount != nil {
		return nil, 0, err
	}
	if err != nil {
		return nil, 0, err
	}
	// Perform any additional business logic or transformations
	return u, count%[1]s, nil
}

func (s *%[2]sUsecase) Create%[1]s(ctx context.Context, name string) (*sqlc.%[1]s, error) {
	u, err := s.%[2]sRepo.Create%[1]s(ctx, name)
	if err != nil {
		return nil, err
	}
	// Perform any additional business logic or transformations
	return u, nil
}

func (s *%[2]sUsecase) Delete%[1]s(ctx context.Context, id string) (err error) {
	err = s.%[2]sRepo.Delete%[1]s(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
`, helper.CapitalizeWord(feat), feat)
	f, err := os.Create("./internal/usecase/" + feat + ".usecase.go")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(sprintf)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Created Usecase to : ./internal/usecase/" + feat + ".usecase.go")
}

func generateHandler(feat string) {
	sprintf := fmt.Sprintf(`package http

import (
	"clean-arch-template/internal/usecase"
	"clean-arch-template/pkg/database/sqlc"
	"clean-arch-template/pkg/helper"
	"clean-arch-template/pkg/response"
	"clean-arch-template/pkg/validation"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type %[1]sResponse struct {
	Success bool
	Data    *sqlc.%[1]s
	Message string
}

var logger = logger2.NewLogger()

type %[1]sHandler struct {
	%[1]sUsecase usecase.%[1]sUsecase
}

func New%[1]sHandler(%[2]sUsecase usecase.%[1]sUsecase) *%[1]sHandler {
	return &%[1]sHandler{
		%[1]sUsecase: %[2]sUsecase,
	}
}

// Get%[1]s is a function to get specific id %[2]s data from database
// @Summary Get %[2]s
// @Description Get %[2]s
// @Tags %[2]ss
// @Accept json
// @Produce json
// @Param id path int true "%[1]s ID"
// @Success 200 {object} %[1]sResponse{}
// @Failure 404 {object} response.Notfound{}
// @Failure 500 {object} response.InternalServer{}
// @Router /%[2]ss/{id} [get]
func (h *%[1]sHandler) Get%[1]s(ec *fiber.Ctx) error {
	id := ec.Params("id")
	%[2]s, err := h.%[1]sUsecase.Get%[1]sByID(ec.UserContext(), id)
	if err != nil {
		logger.Error(err.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Message: "Internal Server Error",
		})
	}

	if %[2]s.ID == 0 {
		ec.Status(fiber.StatusNotFound)
		return ec.JSON(&response.Notfound{
			Message: "Not Found",
		})
	}

	return ec.JSON(&%[1]sResponse{
		Success: true,
		Message: "Berhasil mengambil data",
		Data:    %[2]s,
	})
}

// Delete%[1]s is a function to delete %[2]s data from database
// @Summary Delete %[2]s
// @Description Delete %[2]s
// @Tags %[2]ss
// @Accept json
// @Produce json
// @Param id path int true "%[1]s ID"
// @Success 200 {object} response.SuccessResponse{}
// @Failure 404 {object} response.Notfound{}
// @Failure 500 {object} response.InternalServer{}
// @Router /%[2]ss/{id} [delete]
func (h *%[1]sHandler) Delete%[1]s(ec *fiber.Ctx) error {
	id := ec.Params("id")
	%[2]s, err := h.%[1]sUsecase.Get%[1]sByID(ec.UserContext(), id)
	if err != nil {
		logger.Error(err.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Message: "Internal Server Error",
		})
	}

	if %[2]s.ID == 0 {
		ec.Status(fiber.StatusNotFound)
		return ec.JSON(&response.Notfound{
			Success: false,
			Message: "Not Found",
		})
	}

	errUsecase := h.%[1]sUsecase.Delete%[1]s(ec.UserContext(), id)

	if errUsecase != nil {
		logger.Error(errUsecase.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Success: false,
			Message: "Internal Server Error",
		})
	}

	return ec.JSON(&response.SuccessResponse{
		Success: true,
		Message: "%[1]s Deleted",
	})
}

// List%[1]s is a function to get all %[2]ss data from database
// @Summary Get all %[2]ss
// @Description Get all %[2]ss
// @Tags %[2]ss
// @Accept json
// @Produce json
// @Success 200 {object} response.PaginationResponse{}
// @Failure 500 {object} response.InternalServer{}
// @Router /%[2]ss [get]
func (h *%[1]sHandler) List%[1]ss(ec *fiber.Ctx) error {
	perPage := 15
	currentPage := ec.Query("page")
	currentPageInt, errPage := strconv.Atoi(currentPage)
	baseUrl := ec.BaseURL() + ec.Route().Path

	/**
	Forcing to page 1 if page isn't fulfied converting to int
	*/
	if errPage != nil {
		currentPageInt = 1
	}

	%[2]ss, count, err := h.%[1]sUsecase.List%[1]sPagination(ec.UserContext(), currentPageInt, perPage)

	/**
	Checking error related to usecase
	*/
	if err != nil {
		logger.Error(err.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Success: false,
			Message: "Internal Server Error",
		})
	}

	return ec.JSON(helper.GeneratorPaginationResponse(%[2]ss, int(count), currentPageInt, perPage, baseUrl))
}

// Create%[1]s is a function to create %[2]ss data
// @Summary Create %[2]ss
// @Description Create %[2]ss
// @Tags %[2]ss
// @Accept json
// @Produce json
// @Param name body Validation%[1]s true "%[1]s Name"
// @Success 200 {object} %[1]sResponse{}
// @Failure 500 {object} response.InternalServer{}
// @Router /%[2]ss [post]
func (h *%[1]sHandler) Create%[1]s(ec *fiber.Ctx) error {
	var %[2]s Validation%[1]s
	err := ec.BodyParser(&%[2]s)
	if err != nil {
		logger.Error(err.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Message: err.Error(),
		})
	}

	errors := validation.ValidateStruct(%[2]s)
	if errors != nil {
		return ec.JSON(errors)
	}

	%[2]sModel, err := h.%[1]sUsecase.Create%[1]s(ec.UserContext(), %[2]s.Name)
	if err != nil {
		logger.Error(err.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Message: err.Error(),
		})
	}

	return ec.JSON(&%[1]sResponse{
		Success: true,
		Message: "Berhasil membuat %[2]s",
		Data:    %[2]sModel,
	})
}`, helper.CapitalizeWord(feat), feat)
	f, err := os.Create("./internal/delivery/http/" + feat + ".handler.go")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(sprintf)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Created Usecase to : ./internal/delivery/http/" + feat + ".handler.go")
	fmt.Println("\nNext Steps:")
	fmt.Println("1. Create Migration. (if you did it, skip it)")
	fmt.Println("command: go run cmd/cli/main.go migration create create_" + feat + "_table")
	fmt.Println("2. Adjust column and data type as you need. (if you did it, skip it)")
	fmt.Println("folder: pkg/database/migrations")
	fmt.Println("3. Create all query inside pkg/database/query/query.sql. (if you did it, skip it)")
	fmt.Println("command: -")
	fmt.Println("4. Generate SQLC (if you did it, skip it)")
	fmt.Println("command: sqlc generate")
	fmt.Println("5. Create your own validator inside ./internal/delivery/http/" + feat + ".handler.go")
	fmt.Println("command: -")
}
