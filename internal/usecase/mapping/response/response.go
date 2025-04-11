package response

import (
	"time"

	"github.com/enesanbar/url-shortener/internal/domain"
)

// Presenter is the presenter interface of mapping(s)
type Presenter interface {
	Single(mapping *domain.Mapping) *Response
	Multiple(mapping []*domain.Mapping) []*Response
}

// Response is a mapping response
type Response struct {
	ID        string     `json:"id"`
	Code      string     `json:"code" example:"ZSDASZX"`
	URL       string     `json:"url" example:"https://example.com"`
	ExpiresAt *time.Time `json:"expires_at,omitempty" example:"2021-10-27T14:13:39Z"`
	CreatedAt time.Time  `json:"created_at" example:"2021-10-27T14:13:39.306Z"`
	UpdatedAt time.Time  `json:"updated_at" example:"2021-10-27T14:13:39.306Z"`
}
