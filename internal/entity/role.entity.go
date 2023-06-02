package entity

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Role struct {
	ID        int              `json:"id"`
	Name      pgtype.Text      `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
