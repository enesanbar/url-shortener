package get

import (
	"context"
	"fmt"

	"github.com/enesanbar/url-shortener/internal/domain"

	"github.com/enesanbar/go-service/cache"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
	"go.uber.org/fx"
)

type InteractorCache struct {
	logger    log.Factory
	cache     cache.Cache
	next      Service
	presenter response.Presenter
}

type InteractorCacheParams struct {
	fx.In

	Logger    log.Factory
	Cache     cache.Cache
	Presenter response.Presenter
	Next      Service `name:"interactor"`
}

func NewInteractorCache(p InteractorCacheParams) Service {
	return &InteractorCache{logger: p.Logger, cache: p.Cache, next: p.Next, presenter: p.Presenter}
}

func (i *InteractorCache) Execute(ctx context.Context, request Request) (*domain.Mapping, error) {
	key := fmt.Sprintf("url-shortener.code.%s", request.Code)

	cachedMapping, err := i.cache.Get(ctx, key)
	if err == nil {
		return cachedMapping.(*domain.Mapping), nil
	}

	mapping, err := i.next.Execute(ctx, request)
	if err != nil {
		return nil, err
	}

	i.cache.Set(ctx, key, mapping)

	return mapping, nil
}
