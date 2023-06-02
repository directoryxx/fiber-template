package repository

import (
	"clean-arch-template/pkg/database/sqlc"
	"context"
	"strconv"
)

type RoleRepository interface {
	GetByIDRole(ctx context.Context, id string) (*sqlc.Role, error)
	CreateRole(ctx context.Context, name string) (*sqlc.Role, error)
	ListAllRole(ctx context.Context) (*[]sqlc.Role, error)
	DeleteRole(ctx context.Context, id string) error
	CountRole(ctx context.Context, id string) (count int32, err error)
	CountAllRole(ctx context.Context) (count int32, err error)
	PaginationListRole(ctx context.Context, limit int, offset int) (*[]sqlc.Role, error)
	// other repository methods...
}

type roleRepository struct {
	db *sqlc.Queries
}

func NewRoleRepository(db *sqlc.Queries) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) GetByIDRole(ctx context.Context, id string) (*sqlc.Role, error) {
	idConv, _ := strconv.ParseInt(id, 10, 0)
	role, err := r.db.GetRole(ctx, int32(idConv))
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) ListAllRole(ctx context.Context) (*[]sqlc.Role, error) {
	roleModel, err := r.db.ListRoles(ctx)
	if err != nil {
		return nil, err
	}
	return &roleModel, nil
}

func (r *roleRepository) CreateRole(ctx context.Context, name string) (*sqlc.Role, error) {
	role, err := r.db.CreateRole(ctx, name)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) DeleteRole(ctx context.Context, id string) (err error) {
	idConv, _ := strconv.ParseInt(id, 10, 0)
	err2 := r.db.DeleteRole(ctx, int32(idConv))
	if err2 != nil {
		return err2
	}

	return nil
}

func (r *roleRepository) CountRole(ctx context.Context, id string) (count int32, err error) {
	idConv, _ := strconv.ParseInt(id, 10, 0)
	count, err = r.db.CountRoleByID(ctx, int32(idConv))
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *roleRepository) CountAllRole(ctx context.Context) (count int32, err error) {
	count, err = r.db.CountRoleAll(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *roleRepository) PaginationListRole(ctx context.Context, limit int, offset int) (*[]sqlc.Role, error) {
	args := sqlc.ListRolesPaginationParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	listRole, err := r.db.ListRolesPagination(ctx, args)
	if err != nil {
		return nil, err
	}

	return &listRole, err
}
