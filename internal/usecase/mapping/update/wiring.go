package update

import (
	"context"

	"github.com/enesanbar/url-shortener/internal/domain"
)

type Service interface {
	Execute(ctx context.Context, input *Request) (*domain.Mapping, error)
}
