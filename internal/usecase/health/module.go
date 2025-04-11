package health

import (
	"github.com/enesanbar/go-service/healthchecker"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	healthchecker.AsHealthCheckerProbe(NewChecker),
)
