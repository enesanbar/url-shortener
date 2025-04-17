package deletion

import (
	"context"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/log"
)

// NewDeleteMappingInteractor creates new Interactor with its dependencies
func NewDeleteMappingInteractor(
	logger log.Factory,
	repo Repository,
) Service {
	return &Interactor{
		logger: logger,
		repo:   repo,
	}
}

type Interactor struct {
	logger log.Factory
	repo   Repository
}

// Execute orchestrates the use case
func (i Interactor) Execute(ctx context.Context, code string) error {
	err := i.repo.Delete(ctx, code)
	if err != nil {
		return errors.Error{Op: "DeleteUrlMapping", Err: err}
	}

	return nil
}
