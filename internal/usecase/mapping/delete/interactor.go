package delete

import (
	"context"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
)

// NewDeleteMappingInteractor creates new Interactor with its dependencies
func NewDeleteMappingInteractor(
	logger log.Factory,
	repo Repository,
	presenter response.Presenter,
) *Interactor {
	return &Interactor{
		logger:    logger,
		repo:      repo,
		presenter: presenter,
	}
}

type Interactor struct {
	logger    log.Factory
	repo      Repository
	presenter response.Presenter
}

// Execute orchestrates the use case
func (i Interactor) Execute(ctx context.Context, code string) error {
	err := i.repo.Delete(ctx, code)
	if err != nil {
		return errors.Error{Op: "GetUrlMapping", Err: err}
	}

	return nil
}
