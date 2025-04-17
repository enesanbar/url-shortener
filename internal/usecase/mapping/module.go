package mapping

import (
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/create"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/deletion"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/get"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/getall"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/response"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/update"
	"go.uber.org/fx"
)

var Module = fx.Options(
	create.Module,
	get.Module,
	getall.Module,
	update.Module,
	deletion.Module,
	response.Module,
)
