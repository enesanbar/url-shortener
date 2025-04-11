package repository

import (
	"github.com/enesanbar/url-shortener/internal/adapter/repository/mapping"
	"go.uber.org/fx"
)

var Module = fx.Options(
	mapping.Module,
)
