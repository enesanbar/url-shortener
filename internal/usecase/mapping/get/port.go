package get

import (
	"context"

	"github.com/enesanbar/url-shortener/internal/domain"
)

type Repository interface {
	FindByCode(ctx context.Context, code string) (*domain.Mapping, error)
}
