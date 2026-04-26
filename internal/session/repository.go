package session

import (
	"context"

	"github.com/parfentjev/raweceek/internal/codegen/db"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(queries *db.Queries) *Repository {
	return &Repository{queries}
}

func (r *Repository) GetNextSession(ctx context.Context) (db.Session, error) {
	return r.queries.GetNextSession(ctx)
}

func (r *Repository) CountSessionsThisWeek(ctx context.Context) (int64, error) {
	return r.queries.CountSessionsThisWeek(ctx)
}
