package usecase

import (
	"clean-arch-template/internal/repository"
	"clean-arch-template/pkg/database/sqlc"
	"clean-arch-template/pkg/helper"
	"context"
)

type RoleUsecase interface {
	GetRoleByID(ctx context.Context, id string) (*sqlc.Role, error)
	CreateRole(ctx context.Context, name string) (*sqlc.Role, error)
	ListRole(ctx context.Context) (*[]sqlc.Role, error)
	ListRolePagination(ctx context.Context, page int, page_size int) (*[]sqlc.Role, int32, error)
	DeleteRole(ctx context.Context, id string) error
	// other use case methods...
}

type roleUsecase struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleUsecase {
	return &roleUsecase{roleRepo: roleRepo}
}

func (s *roleUsecase) GetRoleByID(ctx context.Context, id string) (*sqlc.Role, error) {
	u, err := s.roleRepo.GetByIDRole(ctx, id)
	count, errCount := s.roleRepo.CountRole(ctx, id)
	if errCount != nil {
		return nil, err
	}

	if count == 0 {
		roleEmpty := &sqlc.Role{
			ID: 0,
		}
		return roleEmpty, nil
	}

	if err != nil {
		return nil, err
	}

	// Perform any additional business logic or transformations
	return u, nil
}

func (s *roleUsecase) ListRole(ctx context.Context) (*[]sqlc.Role, error) {
	u, err := s.roleRepo.ListAllRole(ctx)
	if err != nil {
		return nil, err
	}
	// Perform any additional business logic or transformations
	return u, nil
}

func (s *roleUsecase) ListRolePagination(ctx context.Context, page int, page_size int) (role *[]sqlc.Role, count int32, err error) {
	offsetCalculate := helper.OffsetCalculator(page, page_size)
	limitCalculate := helper.LimitCalculator(page, page_size)
	u, err := s.roleRepo.PaginationListRole(ctx, limitCalculate, offsetCalculate)
	countRole, errCount := s.roleRepo.CountAllRole(ctx)
	if errCount != nil {
		return nil, 0, err
	}
	if err != nil {
		return nil, 0, err
	}
	// Perform any additional business logic or transformations
	return u, countRole, nil
}

func (s *roleUsecase) CreateRole(ctx context.Context, name string) (*sqlc.Role, error) {
	u, err := s.roleRepo.CreateRole(ctx, name)
	if err != nil {
		return nil, err
	}
	// Perform any additional business logic or transformations
	return u, nil
}

func (s *roleUsecase) DeleteRole(ctx context.Context, id string) (err error) {
	err = s.roleRepo.DeleteRole(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
