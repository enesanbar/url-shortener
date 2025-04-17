package mappings

import (
	"context"
	"net/http"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"github.com/enesanbar/url-shortener/internal/usecase/mapping/deletion"
)

// DeleteMappingsUseCase is the get interface of the service
//
//go:generate mockery -name=DeleteMappingsUseCase
type DeleteMappingsUseCase interface {
	Execute(ctx context.Context, code string) error
}

type DeleteMappingAdapter struct {
	router.BaseHandler
	deleter DeleteMappingsUseCase
}

type DeleteMappingAdapterParams struct {
	fx.In

	BaseHandler router.BaseHandler
	Interactor  deletion.Service `name:"producer"`
}

func NewDeleteMappingAdapter(p DeleteMappingAdapterParams) *DeleteMappingAdapter {
	return &DeleteMappingAdapter{
		BaseHandler: p.BaseHandler,
		deleter:     p.Interactor,
	}
}

// Handle godoc
//
//	@Summary		Deletes an existing mapping
//	@Description	Deletes an existing mapping
//	@Tags			mappings
//	@Param			code	path		string	true	"Short Code"
//	@Success		204		{string}	string	""		desc
//	@Router			/mappings/{code} [delete]
func (h DeleteMappingAdapter) Handle(c echo.Context) error {
	code := c.Param("code")

	err := h.deleter.Execute(c.Request().Context(), code)
	if err != nil {
		return h.NewError(c, errors.Error{Op: "DeleteMappingHandler.Handle", Err: err})
	}

	return c.NoContent(http.StatusNoContent)
}
