package mappings

import (
	"context"
	"net/http"

	"go.uber.org/fx"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/router"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/getall"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
	"github.com/labstack/echo/v4"
)

// GetMappingsUsecase is the port to the getall use case
type GetMappingsUsecase interface {
	Execute(ctx context.Context, request *getall.Request) (*router.PagedResponse, error)
}

type GetMappingsAdapter struct {
	router.BaseHandler
	getter    GetMappingsUsecase
	presenter response.Presenter
}

type GetMappingsAdapterParams struct {
	fx.In

	BaseHandler        router.BaseHandler
	Presenter          response.Presenter
	GetMappingsUseCase GetMappingsUsecase
}

func NewGetMappingsAdapter(p GetMappingsAdapterParams) *GetMappingsAdapter {
	return &GetMappingsAdapter{
		BaseHandler: p.BaseHandler,
		getter:      p.GetMappingsUseCase,
		presenter:   p.Presenter,
	}
}

// Handle godoc
//
//	@Summary		Get URL Mappings
//	@Description	Get all URL mappings
//	@Tags			mappings
//	@Param			page		query		int																			false	"Page Number"	minimum(1)			default(1)
//	@Param			pageSize	query		int																			false	"Page Size"		minimum(1)			maximum(64)	default(16)
//	@Param			sortOrder	query		string																		false	"Sort Order"	Enums(asc, desc)	default(desc)
//	@Param			sortBy		query		string																		false	"Sort Field"
//	@Param			codeQuery	query		string																		false	"Search By Code"
//	@Param			urlQuery	query		string																		false	"Search By URL"
//	@Success		200			{object}	router.ApiResponse{data=router.PagedResponse{items=[]response.Response}}	"desc"
//	@Router			/mappings [get]
func (h GetMappingsAdapter) Handle(c echo.Context) error {
	page := c.QueryParam("page")
	pageSize := c.QueryParam("pageSize")
	sortOrder := c.QueryParam("sortOrder")
	sortBy := c.QueryParam("sortBy")
	codeQuery := c.QueryParam("codeQuery")
	urlQuery := c.QueryParam("urlQuery")
	request := getall.NewRequest(page, pageSize, sortBy, sortOrder, codeQuery, urlQuery)

	responseObject, err := h.getter.Execute(c.Request().Context(), request)
	if err != nil {
		return h.NewError(c, errors.Error{Op: "GetMappingAdapter.Handle", Err: err})
	}

	return h.NewSuccess(c, responseObject, http.StatusOK)
}
