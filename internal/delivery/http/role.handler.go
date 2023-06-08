package http

import (
	"clean-arch-template/internal/usecase"
	"clean-arch-template/pkg/database/sqlc"
	"clean-arch-template/pkg/response"
	"clean-arch-template/pkg/utils"
	"clean-arch-template/pkg/validation"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type RoleResponse struct {
	Success bool
	Data    *sqlc.Role
	Message string
}

type RoleHandler struct {
	RoleUsecase usecase.RoleUsecase
}

func NewRoleHandler(roleUsecase usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{
		RoleUsecase: roleUsecase,
	}
}

// GetRole is a function to get specific id role data from database
// @Summary Get role
// @Description Get role
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} RoleResponse{}
// @Failure 404 {object} response.Notfound{}
// @Failure 500 {object} response.InternalServer{}
// @Router /roles/{id} [get]
func (h *RoleHandler) GetRole(ec *fiber.Ctx) error {
	id := ec.Params("id")
	role, err := h.RoleUsecase.GetRoleByID(ec.UserContext(), id)
	if err != nil {
		newLogger.Error(err.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Message: "Internal Server Error",
		})
	}

	if role.ID == 0 {
		ec.Status(fiber.StatusNotFound)
		return ec.JSON(&response.Notfound{
			Message: "Not Found",
		})
	}

	return ec.JSON(&RoleResponse{
		Success: true,
		Message: "Berhasil mengambil data",
		Data:    role,
	})
}

// DeleteRole is a function to delete role data from database
// @Summary Delete role
// @Description Delete role
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} response.SuccessResponse{}
// @Failure 404 {object} response.Notfound{}
// @Failure 500 {object} response.InternalServer{}
// @Router /roles/{id} [delete]
func (h *RoleHandler) DeleteRole(ec *fiber.Ctx) error {
	id := ec.Params("id")
	role, err := h.RoleUsecase.GetRoleByID(ec.UserContext(), id)
	if err != nil {
		newLogger.Error(err.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Message: "Internal Server Error",
		})
	}

	if role.ID == 0 {
		ec.Status(fiber.StatusNotFound)
		return ec.JSON(&response.Notfound{
			Success: false,
			Message: "Not Found",
		})
	}

	errUsecase := h.RoleUsecase.DeleteRole(ec.UserContext(), id)

	if errUsecase != nil {
		newLogger.Error(errUsecase.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Success: false,
			Message: "Internal Server Error",
		})
	}

	return ec.JSON(&response.SuccessResponse{
		Success: true,
		Message: "Role Deleted",
	})
}

// ListRole is a function to get all roles data from database
// @Summary Get all roles
// @Description Get all roles
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} response.PaginationResponse{}
// @Failure 500 {object} response.InternalServer{}
// @Router /roles [get]
func (h *RoleHandler) ListRoles(ec *fiber.Ctx) error {
	span := sentry.StartSpan(ec.UserContext(), "handler",
		sentry.WithTransactionName(ec.Method()+" "+ec.OriginalURL()))

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

	roles, count, err := h.RoleUsecase.ListRolePagination(ec.UserContext(), currentPageInt, perPage)

	span.Finish()

	/**
	Checking error related to usecase
	*/
	if err != nil {
		newLogger.Error(err.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Success: false,
			Message: "Internal Server Error",
		})
	}

	return ec.JSON(utils.GeneratorPaginationResponse(roles, int(count), currentPageInt, perPage, baseUrl))
}

// CreateRole is a function to create roles data
// @Summary Create roles
// @Description Create roles
// @Tags roles
// @Accept json
// @Produce json
// @Param name body ValidationRole true "Role Name"
// @Success 200 {object} RoleResponse{}
// @Failure 500 {object} response.InternalServer{}
// @Router /roles [post]
func (h *RoleHandler) CreateRole(ec *fiber.Ctx) error {
	var role ValidationRole
	err := ec.BodyParser(&role)
	if err != nil {
		newLogger.Error(err.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Message: err.Error(),
		})
	}

	errors := validation.ValidateStruct(role)
	if errors != nil {
		return ec.JSON(errors)
	}

	roleModel, err := h.RoleUsecase.CreateRole(ec.UserContext(), role.Name)
	if err != nil {
		newLogger.Error(err.Error())
		ec.Status(fiber.StatusInternalServerError)
		return ec.JSON(&response.InternalServer{
			Message: err.Error(),
		})
	}

	ec.Status(fiber.StatusCreated)
	return ec.JSON(&RoleResponse{
		Success: true,
		Message: "Berhasil membuat role",
		Data:    roleModel,
	})
}

type ValidationRole struct {
	Name string `json:"name" validate:"required"`
}
