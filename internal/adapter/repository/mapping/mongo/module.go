package mongo

import (
	"github.com/enesanbar/go-service/persistance/mongodb"
	"github.com/enesanbar/go-service/wiring"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/create"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/delete"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/get"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/getall"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/update"
	"github.com/enesanbar/url-shortener/internal/usecase/redirect"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	bindings,
)

var factories = fx.Provide(
	fx.Annotated{
		Name:   "default",
		Target: NewMongoCollection,
	},
	fx.Annotated{
		Name:   "default",
		Target: NewDatabase,
	},
	// add client to the connections group to close at app stop
	fx.Annotated{
		Group: "connections",
		Target: func(p DBParams) wiring.Connection {
			return p.Conn
		},
	},
	fx.Annotated{
		Name:   "default",
		Target: NewMongoClient,
	},
	fx.Annotated{
		Name:   "default",
		Target: NewConfig,
	},
	mongodb.NewConnector,
)
var bindings = fx.Provide(
	fx.Annotate(
		NewMappingMongoAdapter,
		fx.As(new(create.Repository)),
		fx.As(new(delete.Repository)),
		fx.As(new(get.Repository)),
		fx.As(new(getall.Repository)),
		fx.As(new(update.Repository)),
		fx.As(new(redirect.Repository)),
	),
)
