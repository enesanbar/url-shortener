package deletion

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	fx.Annotated{
		Name:   "interactor",
		Target: NewDeleteMappingInteractor,
	},
	fx.Annotated{
		Name:   "producer",
		Target: NewDeleteMappingInteractorProducer,
	},
)
