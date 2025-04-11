package create

import (
	"context"

	"github.com/enesanbar/url-shortener/internal/domain"
)

//go:generate mockery -name=Repository
type Repository interface {
	Store(ctx context.Context, mapping *domain.Mapping) (*domain.Mapping, error)
	FindByCode(ctx context.Context, code string) (*domain.Mapping, error)
}
