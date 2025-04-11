package redirect

import (
	"context"
	"time"

	"github.com/enesanbar/url-shortener/internal/domain"

	"github.com/enesanbar/go-service/config"

	"github.com/enesanbar/go-service/log"
	"go.uber.org/fx"
)

type RetrierConfig struct {
	WaitInterval time.Duration
	RetryCount   int
}

func NewRetrierConfig(cfg config.Config) *RetrierConfig {
	// TODO: Set default values if not provided
	return &RetrierConfig{
		WaitInterval: time.Duration(cfg.GetInt("usecases.redirect.retrier.wait-interval")) * time.Millisecond,
		RetryCount:   cfg.GetInt("usecases.redirect.retrier.retry-count"),
	}
}

type Retrier struct {
	log  log.Factory
	next Interactor
	cfg  *RetrierConfig
}

type RetrierParams struct {
	fx.In

	Log           log.Factory
	Next          Interactor `name:"interactor"`
	RetrierConfig *RetrierConfig
}

func NewRetrier(p RetrierParams) Interactor {
	return &Retrier{log: p.Log, next: p.Next, cfg: p.RetrierConfig}
}

func (r *Retrier) Execute(ctx context.Context, request *Request) (*domain.Mapping, error) {
	var data *domain.Mapping
	var err error
	for retry := 1; retry <= r.cfg.RetryCount; retry++ {
		if data, err = r.next.Execute(ctx, request); err == nil {
			return data, nil
		} else if retry == r.cfg.RetryCount {
			r.log.For(ctx).Infof("failed to fetch, tried %d times", retry)
			return nil, err
		}

		r.log.For(ctx).Infof("attempt %d failed, retrying in %s", retry, r.cfg.WaitInterval)

		select {
		case <-time.After(r.cfg.WaitInterval):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	return data, err
}
