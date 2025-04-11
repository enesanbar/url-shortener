package adapter

import (
	"github.com/enesanbar/url-shortener/internal/adapter/cache"
	"github.com/enesanbar/url-shortener/internal/adapter/repository"
	"github.com/enesanbar/url-shortener/internal/adapter/rest"
	"github.com/enesanbar/url-shortener/internal/usecase"
	"go.uber.org/fx"

	_ "github.com/enesanbar/url-shortener/docs"
)

var Module = fx.Options(
	repository.Module,
	rest.Module,
	cache.Module,
	usecase.Module,
)
