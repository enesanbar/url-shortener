package mapping

import (
	"github.com/enesanbar/url-shortener/internal/adapter/repository/mapping/mongo"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// inject different datasource based on config later
	mongo.Module,
)
