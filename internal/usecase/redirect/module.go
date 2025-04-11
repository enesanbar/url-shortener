package redirect

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	fx.Annotated{
		Name:   "interactor",
		Target: NewInteractor,
	},
	fx.Annotated{
		Name:   "retrier",
		Target: NewRetrier,
	},
	NewRetrierConfig,
)
