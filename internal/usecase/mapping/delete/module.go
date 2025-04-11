package delete

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	NewDeleteMappingInteractor,
)
