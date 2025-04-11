package redirect

import (
	"go.uber.org/fx"

	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/url-shortener/internal/usecase/redirect"
)

type Adapter struct {
	getter redirect.Interactor
	logger log.Factory
}

type AdapterParams struct {
	fx.In

	Logger     log.Factory
	Interactor redirect.Interactor `name:"retrier"`
}

func New(p AdapterParams) *Adapter {
	return &Adapter{getter: p.Interactor, logger: p.Logger}
}
