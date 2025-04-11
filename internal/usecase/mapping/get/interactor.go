package get

import (
	"context"

	"github.com/enesanbar/go-service/log"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/url-shortener/internal/domain"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
)

type Interactor struct {
	logger    log.Factory
	repo      Repository
	presenter response.Presenter
}

// NewGetMappingInteractor creates new Interactor with its dependencies
func NewGetMappingInteractor(p Params) Service {
	return &Interactor{
		logger:    p.Logger,
		repo:      p.Repo,
		presenter: p.Presenter,
	}
}

// Execute orchestrates the use case
func (i Interactor) Execute(ctx context.Context, request Request) (*domain.Mapping, error) {
	result, err := i.repo.FindByCode(ctx, request.Code)
	if err != nil {
		return &domain.Mapping{}, errors.Error{Op: "GetUrlMapping", Err: err}
	}

	return result, nil
}
