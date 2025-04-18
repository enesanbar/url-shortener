package update

import (
	"context"

	"github.com/enesanbar/url-shortener/internal/domain"
)

type Repository interface {
	FindByCode(ctx context.Context, code string) (*domain.Mapping, error)
	Update(ctx context.Context, mapping *domain.Mapping) (*domain.Mapping, error)
}
