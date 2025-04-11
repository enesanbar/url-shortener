package cache

import (
	"github.com/enesanbar/go-service/cache/inmemory"
	"github.com/enesanbar/go-service/config"
	"github.com/enesanbar/go-service/log"
)

func NewConfig(cfg config.Config, logger log.Factory) (*inmemory.Config, error) {
	return inmemory.New(cfg, logger, "cache.inmemory")
}
