package update

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	fx.Annotated{
		Name:   "interactor",
		Target: NewUpdateMappingInteractor,
	},
	fx.Annotated{
		Name:   "producer",
		Target: NewUpdateMappingInteractorProducer,
	},
)
