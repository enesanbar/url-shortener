package mappings

import (
	"context"
	"net/http"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/router"
	"github.com/labstack/echo/v4"

	"github.com/enesanbar/url-shortener/internal/domain"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/create"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
)

// CreateMappingUseCase is the port to the create mapping use case
//
//go:generate mockery --name=CreateMappingUseCase
type CreateMappingUseCase interface {
	Execute(context.Context, *create.Request) (*domain.Mapping, error)
}

type CreateMappingAdapter struct {
	router.BaseHandler
	creator   CreateMappingUseCase
	presenter response.Presenter
	logger    log.Factory
}

func NewCreateMappingAdapter(
	baseHandler router.BaseHandler,
	c CreateMappingUseCase,
	presenter response.Presenter,
	l log.Factory,
) *CreateMappingAdapter {
	return &CreateMappingAdapter{
		BaseHandler: baseHandler,
		creator:     c,
		presenter:   presenter,
		logger:      l,
	}
}

// Handle godoc
//	@Summary		Create or Generate URL Mapping
//	@Description	If 'code' parameter is not supplied, one will be generated
//	@Tags			mappings
//	@Param			mapping	body		create.Request								true	"Mapping Request"
//	@Success		201		{object}	router.ApiResponse{data=response.Response}	"desc"
//	@Router			/mappings [post]
func (h CreateMappingAdapter) Handle(c echo.Context) error {
	var request create.Request
	err := h.DecodeRequest(c, &request)
	if err != nil {
		return h.NewError(c, errors.Error{Op: "CreateMappingAdapter", Err: err})
	}

	responseObject, err := h.creator.Execute(c.Request().Context(), &request)
	if err != nil {
		return h.NewError(c, errors.Error{Op: "CreateMappingAdapter", Err: err})
	}

	return h.NewSuccess(c, h.presenter.Single(responseObject), http.StatusCreated)
}
