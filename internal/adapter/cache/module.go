package cache

import (
	"github.com/enesanbar/go-service/cache"
	"github.com/enesanbar/go-service/cache/inmemory"
	"github.com/enesanbar/go-service/cache/metrics"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"cache",
	fx.Provide(
		NewConfig,
		metrics.NewInstrumentor,
		fx.Annotate(
			inmemory.NewInMemoryCache,
			fx.As(new(cache.Cache)),
		),
	),
)
