package usecase

import (
	"github.com/enesanbar/url-shortener/internal/usecase/health"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping"
	"github.com/enesanbar/url-shortener/internal/usecase/redirect"
	"go.uber.org/fx"
)

var Module = fx.Options(
	mapping.Module,
	redirect.Module,
	health.Module,
)
