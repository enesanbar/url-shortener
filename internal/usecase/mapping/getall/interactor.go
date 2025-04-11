package getall

import (
	"context"

	"github.com/enesanbar/go-service/router"

	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
)

// NewGetMappingsInteractor creates new Interactor with its dependencies
func NewGetMappingsInteractor(
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
func (i Interactor) Execute(ctx context.Context, request *Request) (*router.PagedResponse, error) {
	return i.repo.FindAll(ctx, request)
}
