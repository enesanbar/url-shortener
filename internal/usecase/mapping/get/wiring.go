package get

import (
	"context"

	"github.com/enesanbar/url-shortener/internal/domain"

	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Logger    log.Factory
	Repo      Repository
	Presenter response.Presenter
}

type Service interface {
	Execute(ctx context.Context, request Request) (*domain.Mapping, error)
}
