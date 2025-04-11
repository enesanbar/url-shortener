package redirect

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	New,
	fx.Annotate(
		RegisterRoutes,
		fx.ResultTags(`group:"routes"`),
	),
)
