package redirect

import (
	"context"

	"github.com/enesanbar/url-shortener/internal/usecase/mapping/get"

	"github.com/enesanbar/url-shortener/internal/domain"

	"github.com/enesanbar/go-service/errors"

	"go.uber.org/fx"

	"github.com/enesanbar/go-service/log"
)

type Params struct {
	fx.In

	Logger        log.Factory
	GetInteractor get.Service `name:"cache"`
	Config        *domain.AppConfig
}

// NewInteractor creates new interactor with its dependencies
func NewInteractor(p Params) Interactor {
	return &interactor{
		logger: p.Logger,
		getter: p.GetInteractor,
		cfg:    p.Config,
	}
}

type interactor struct {
	logger log.Factory
	getter get.Service
	cfg    *domain.AppConfig
}

// Execute orchestrates the use case
func (a *interactor) Execute(ctx context.Context, input *Request) (*domain.Mapping, error) {
	mapping, err := a.getter.Execute(ctx, get.Request{Code: input.Code})
	if err != nil {
		return nil, errors.Error{Err: err}
	}

	if mapping.IsExpired() {
		mapping.URL = a.cfg.DefaultRedirectURL
	}

	return mapping, nil
}
