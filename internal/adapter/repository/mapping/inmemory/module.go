package inmemory

import (
	"github.com/enesanbar/url-shortener/internal/domain"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/create"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/delete"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/get"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/update"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	bindings,
)

var factories = fx.Provide(
	func() map[string]*domain.Mapping {
		return make(map[string]*domain.Mapping)
	},
	NewMappingInmemoryAdapter,
)

var bindings = fx.Provide(
	func(a *MappingInmemoryAdapter) create.Repository { return a },
	func(a *MappingInmemoryAdapter) delete.Repository { return a },
	func(a *MappingInmemoryAdapter) get.Repository { return a },
	//func(a *MappingInmemoryAdapter) getall.Repository { return a },
	func(a *MappingInmemoryAdapter) update.Repository { return a },
)
