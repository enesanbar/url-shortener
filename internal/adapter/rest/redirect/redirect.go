package redirect

import (
	"errors"
	"net/http"

	"github.com/enesanbar/url-shortener/internal/domain"
	"go.uber.org/zap"

	"github.com/enesanbar/url-shortener/internal/usecase/redirect"
	"github.com/labstack/echo/v4"
)

func (h *Adapter) RedirectShortUrl(c echo.Context) error {
	code := c.Param(PathParamCode)
	url, err := h.getter.Execute(c.Request().Context(), &redirect.Request{Code: code})

	if err != nil {
		if !errors.Is(err, domain.ErrMappingNotFound) {
			h.logger.For(c.Request().Context()).
				With(zap.String("code", code)).
				With(zap.Error(err)).
				Error("an error occurred")
		}

		return c.String(http.StatusNotFound, "not found")
	}

	return c.Redirect(http.StatusMovedPermanently, url.URL)
}
