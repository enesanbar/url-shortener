package main

import (
	"github.com/enesanbar/go-service/service"
	"github.com/enesanbar/url-shortener/internal/adapter"
	"github.com/enesanbar/url-shortener/internal/domain"
	"github.com/enesanbar/url-shortener/internal/validators"
)

func main() {
	service.New("url-shortener").
		WithRestAdapter().
		WithModules(
			adapter.Module,
			validators.Module,
		).
		WithConstructor(domain.NewAppConfig).
		Build().Run()
}
