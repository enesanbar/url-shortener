package redirect

import (
	"github.com/enesanbar/go-service/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	New,
	router.AsRoute(RegisterRoutes),
)
