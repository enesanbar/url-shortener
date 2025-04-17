package update

import (
	"context"
	"time"

	"github.com/enesanbar/url-shortener/internal/domain"

	"go.uber.org/fx"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/validation"
)

type Interactor struct {
	logger    log.Factory
	repo      Repository
	validator validation.Validator
}

type Params struct {
	fx.In

	Logger    log.Factory
	Repo      Repository
	Validator validation.Validator `name:"go_playground"`
}

// NewUpdateMappingInteractor creates new Interactor with its dependencies
func NewUpdateMappingInteractor(p Params) Service {
	return &Interactor{
		logger:    p.Logger,
		repo:      p.Repo,
		validator: p.Validator,
	}
}

// Execute orchestrates the use case
func (a Interactor) Execute(ctx context.Context, input *Request) (*domain.Mapping, error) {
	// validation may be moved to router.BaseHandler
	err := a.validator.Validate(input)

	if err != nil {
		return nil, errors.Error{
			Code:    errors.EINVALID,
			Message: "validation error",
			Err:     err,
			Data:    a.validator.Messages(err),
		}
	}

	storedMapping, err := a.repo.FindByCode(ctx, input.Code)
	if err != nil {
		return nil, errors.Error{
			Err:     err,
			Message: "mapping not found",
		}
	}

	expiresAt, err := storedMapping.NewDateFromLayout("2006-01-02 15:04:05", input.ExpiresAt)
	if err != nil {
		return nil, err
	}

	storedMapping.URL = input.URL
	storedMapping.ExpiresAt = expiresAt
	storedMapping.UpdatedAt = time.Now().Truncate(time.Millisecond)

	result, err := a.repo.Update(ctx, storedMapping)
	if err != nil {
		return nil, err
	}

	return result, err
}
