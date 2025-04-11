package cache

import (
	"github.com/enesanbar/go-service/cache"
	"github.com/enesanbar/go-service/cache/inmemory"
	"github.com/enesanbar/go-service/cache/metrics"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	bindings,
)

var factories = fx.Provide(
	NewConfig,
	inmemory.NewInMemoryCache,
	metrics.NewInstrumentor,
)

var bindings = fx.Provide(
	func(cache *inmemory.Cache) cache.Cache { return cache },
)
