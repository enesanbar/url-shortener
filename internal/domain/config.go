package domain

import (
	"github.com/enesanbar/go-service/config"
)

type AppConfig struct {
	DefaultRedirectURL string
}

func NewAppConfig(cfg config.Config) (*AppConfig, error) {
	defaultRedirectUrl := cfg.GetString("default-redirect")
	if defaultRedirectUrl == "" {
		defaultRedirectUrl = "https://www.google.com"
	}

	return &AppConfig{
		DefaultRedirectURL: defaultRedirectUrl,
	}, nil
}
