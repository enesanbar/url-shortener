package create

import (
	"context"
	"time"

	"github.com/teris-io/shortid"

	"go.uber.org/fx"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/validation"
	"github.com/enesanbar/url-shortener/internal/domain"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Logger    log.Factory
	Repo      Repository
	Presenter response.Presenter
	Validator validation.Validator `name:"go_playground"`
}

// NewCreateMappingInteractor creates new Interactor with its dependencies
func NewCreateMappingInteractor(p Params) *Interactor {
	return &Interactor{
		logger:    p.Logger,
		repo:      p.Repo,
		presenter: p.Presenter,
		validator: p.Validator,
	}
}

type Interactor struct {
	logger    log.Factory
	repo      Repository
	presenter response.Presenter
	validator validation.Validator
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

	a.logger.For(ctx).
		With(zap.String("code", input.Code), zap.String("url", input.URL)).
		Info("creating mapping")

	// if user did not provide code, generate one
	if input.Code == "" {
		input.Code = shortid.MustGenerate()
	}

	_, err = a.repo.FindByCode(ctx, input.Code)
	if err == nil {
		return &domain.Mapping{}, errors.Error{
			Code:    errors.ECONFLICT,
			Message: "code already exists",
			Op:      "shortener.ServiceImpl > Handle",
			Err:     err,
		}
	}

	var newMapping domain.Mapping
	expiresAt, err := newMapping.NewDateFromLayout("2006-01-02 15:04:05", input.ExpiresAt)
	if err != nil {
		return nil, err
	}

	newMapping.Code = input.Code
	newMapping.URL = input.URL
	newMapping.ExpiresAt = expiresAt
	newMapping.CreatedAt = time.Now().Truncate(time.Millisecond)
	newMapping.UpdatedAt = time.Now().Truncate(time.Millisecond)

	storedMapping, err := a.repo.Store(ctx, &newMapping)
	if err != nil {
		return nil, errors.Error{Err: err}
	}

	return storedMapping, nil
}
