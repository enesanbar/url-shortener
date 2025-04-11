package rest

import (
	"github.com/enesanbar/url-shortener/internal/adapter/rest/mappings"
	"github.com/enesanbar/url-shortener/internal/adapter/rest/redirect"
	"go.uber.org/fx"
)

var Module = fx.Options(
	mappings.Module,
	redirect.Module,
)
