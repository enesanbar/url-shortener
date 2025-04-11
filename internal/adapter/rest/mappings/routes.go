package mappings

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"

	"github.com/enesanbar/go-service/router"
)

var (
	Base          = "/api/mappings"
	PathParamCode = "code"

	Create       = ""
	GetAll       = ""
	GetByCode    = fmt.Sprintf("/:%s", PathParamCode)
	UpdateByCode = fmt.Sprintf("/:%s", PathParamCode)
	DeleteByCode = GetByCode
)

type Params struct {
	fx.In

	Create *CreateMappingAdapter
	Get    *GetMappingAdapter
	GetAll *GetMappingsAdapter
	Update *UpdateMappingAdapter
	Delete *DeleteMappingAdapter
}

func RegisterRoutes(p Params) router.RouteConfig {
	return router.RouteConfig{
		Path: Base,
		Router: func(group *echo.Group) {
			group.Use(middleware.CORS())
			group.POST(Create, p.Create.Handle)
			group.GET(GetByCode, p.Get.Handle)
			group.GET(GetAll, p.GetAll.Handle)
			group.PATCH(UpdateByCode, p.Update.Handle)
			group.DELETE(DeleteByCode, p.Delete.Handle)
		},
	}
}
