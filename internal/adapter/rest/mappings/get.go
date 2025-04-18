package mappings

import (
	"context"
	"net/http"

	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"

	"github.com/enesanbar/url-shortener/internal/domain"

	"go.uber.org/fx"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/router"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/get"
	"github.com/labstack/echo/v4"
)

// GetMappingUsecase is the port to the get use case
type GetMappingUsecase interface {
	Execute(ctx context.Context, request get.Request) (*domain.Mapping, error)
}

type GetMappingAdapter struct {
	router.BaseHandler
	getter    GetMappingUsecase
	presenter response.Presenter
}

type GetMappingAdapterParams struct {
	fx.In

	BaseHandler router.BaseHandler
	Presenter   response.Presenter
	Interactor  get.Service `name:"cache"`
}

func NewGetMappingAdapter(p GetMappingAdapterParams) *GetMappingAdapter {
	return &GetMappingAdapter{
		BaseHandler: p.BaseHandler,
		getter:      p.Interactor,
		presenter:   p.Presenter,
	}
}

// Handle godoc
//
//	@Summary		Get URL Mapping
//	@Description	Get a single URL mapping
//	@Tags			mappings
//	@Param			code	path		string										true	"Short Code"
//	@Success		200		{object}	router.ApiResponse{data=response.Response}	"desc"
//	@Failure		404		{object}	router.ApiResponse							"desc"
//	@Router			/mappings/{code} [get]
func (h GetMappingAdapter) Handle(c echo.Context) error {
	code := c.Param("code")
	mapping, err := h.getter.Execute(c.Request().Context(), get.Request{Code: code})
	if err != nil {
		return h.NewError(c, errors.Error{Op: "GetMappingAdapter", Err: err})
	}

	return h.NewSuccess(c, h.presenter.Single(mapping), http.StatusOK)
}
