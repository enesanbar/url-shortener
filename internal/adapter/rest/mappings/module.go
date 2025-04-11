package mappings

import (
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/create"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/delete"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/getall"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/update"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	bindings,
)

var factories = fx.Provide(
	NewCreateMappingAdapter,
	NewGetMappingAdapter,
	NewGetMappingsAdapter,
	NewDeleteMappingAdapter,
	NewUpdateMappingAdapter,
	fx.Annotate(
		RegisterRoutes,
		fx.ResultTags(`group:"routes"`),
	),
)

var bindings = fx.Provide(
	func(service *create.Interactor) CreateMappingUseCase { return service },
	func(service *getall.Interactor) GetMappingsUsecase { return service },
	func(service *update.Interactor) UpdateMappingUseCase { return service },
	func(service *delete.Interactor) DeleteMappingsUseCase { return service },
)
