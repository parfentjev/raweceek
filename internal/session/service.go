package session

import (
	"context"

	"github.com/parfentjev/raweceek/internal/codegen/db"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository}
}

func (s *Service) GetNextSession(ctx context.Context) (db.Session, error) {
	return s.repository.GetNextSession(ctx)
}
