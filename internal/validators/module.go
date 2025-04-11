package validators

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotated{
		Group:  "validators",
		Target: NewIsURLValidator,
	},
)
