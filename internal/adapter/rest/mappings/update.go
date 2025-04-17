package mappings

import (
	"context"
	"net/http"

	"go.uber.org/fx"

	"github.com/enesanbar/url-shortener/internal/domain"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/router"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/update"
	"github.com/labstack/echo/v4"
)

// UpdateMappingUseCase is the port to the update use case
//
//go:generate mockery -name=UpdateMappingUseCase
type UpdateMappingUseCase interface {
	Execute(context.Context, *update.Request) (*domain.Mapping, error)
}

type UpdateMappingAdapter struct {
	router.BaseHandler
	updater   UpdateMappingUseCase
	presenter response.Presenter
	logger    log.Factory
}

type UpdateMappingAdapterParams struct {
	fx.In

	BaseHandler          router.BaseHandler
	UpdateMappingUseCase update.Service `name:"producer"`
	Presenter            response.Presenter
	Logger               log.Factory
}

func NewUpdateMappingAdapter(p UpdateMappingAdapterParams) *UpdateMappingAdapter {
	return &UpdateMappingAdapter{
		BaseHandler: p.BaseHandler,
		updater:     p.UpdateMappingUseCase,
		presenter:   p.Presenter,
		logger:      p.Logger,
	}
}

// Handle godoc
//
//	@Summary		Update existing URL Mapping
//	@Description	Update existing URL Mapping
//	@Tags			mappings
//	@Param			code	path		string										true	"Short Code"
//	@Param			mapping	body		update.Request								true	"Update Request"
//	@Success		202		{object}	router.ApiResponse{data=response.Response}	"desc"
//	@Router			/mappings/{code} [patch]
func (h UpdateMappingAdapter) Handle(c echo.Context) error {
	var request update.Request
	err := h.DecodeRequest(c, &request)
	if err != nil {
		return h.NewError(c, errors.Error{Op: "CreateMappingAdapter", Err: err})
	}
	request.Code = c.Param("code")

	responseObject, err := h.updater.Execute(c.Request().Context(), &request)
	if err != nil {
		return h.NewError(c, errors.Error{Op: "CreateMappingAdapter", Err: err})
	}

	return h.NewSuccess(c, h.presenter.Single(responseObject), http.StatusAccepted)
}
