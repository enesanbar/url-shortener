package create

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	fx.Annotated{
		Name:   "interactor",
		Target: NewCreateMappingInteractor,
	},
	fx.Annotated{
		Name:   "producer",
		Target: NewCreateMappingInteractorProducer,
	},
)
