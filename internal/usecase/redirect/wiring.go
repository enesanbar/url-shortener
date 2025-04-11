package redirect

import (
	"context"

	"github.com/enesanbar/url-shortener/internal/domain"
)

type Interactor interface {
	Execute(ctx context.Context, request *Request) (*domain.Mapping, error)
}
