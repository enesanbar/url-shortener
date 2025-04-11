package response

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	bindings,
)

var factories = fx.Provide(
	NewMappingPresenter,
)
var bindings = fx.Provide(
	func(presenter *MappingPresenter) Presenter {
		return presenter
	},
)
