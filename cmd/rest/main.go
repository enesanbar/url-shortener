package main

import (
	"github.com/enesanbar/go-service/service"
	"github.com/enesanbar/url-shortener/internal/adapter"
	"github.com/enesanbar/url-shortener/internal/adapter/events"
	"github.com/enesanbar/url-shortener/internal/adapter/grpc/mappings"
	"github.com/enesanbar/url-shortener/internal/domain"
	"github.com/enesanbar/url-shortener/internal/validators"
)

func main() {
	service.New("url-shortener").
		WithRestAdapter().
		WithConsumer().
		WithModules(
			adapter.Module,
			validators.Module,
			events.Module,
			mappings.Module, // for now, we are using the same service for both gRPC and REST, move it to a separate cmd later
		).
		WithConstructor(domain.NewAppConfig).
		Build().Run()
}
