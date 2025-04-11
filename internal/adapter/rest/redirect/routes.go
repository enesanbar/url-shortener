package redirect

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/enesanbar/go-service/router"
	"go.uber.org/fx"
)

var (
	Base          = "/redirect"
	PathParamCode = "code"

	Redirect = fmt.Sprintf("/:%s", PathParamCode)
)

type Params struct {
	fx.In

	RedirectHandler *Adapter
}

func RegisterRoutes(p Params) router.RouteConfig {
	return router.RouteConfig{
		Path: Base,
		Router: func(group *echo.Group) {
			group.GET(Redirect, p.RedirectHandler.RedirectShortUrl)
		},
	}
}
