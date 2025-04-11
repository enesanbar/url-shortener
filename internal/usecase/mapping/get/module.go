package get

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	fx.Annotated{
		Name:   "interactor",
		Target: NewGetMappingInteractor,
	},
	fx.Annotated{
		Name:   "cache",
		Target: NewInteractorCache,
	},
)
