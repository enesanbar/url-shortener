package mappings

import (
	"context"
	"net/http"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/router"
	"github.com/labstack/echo/v4"
)

//go:generate mockery -name=DeleteMappingsUseCase
// DeleteMappingsUseCase is the get interface of the service
type DeleteMappingsUseCase interface {
	Execute(ctx context.Context, code string) error
}

type DeleteMappingAdapter struct {
	router.BaseHandler
	deleter DeleteMappingsUseCase
}

func NewDeleteMappingAdapter(baseHandler router.BaseHandler, rs DeleteMappingsUseCase) *DeleteMappingAdapter {
	return &DeleteMappingAdapter{BaseHandler: baseHandler, deleter: rs}
}

// Handle godoc
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
