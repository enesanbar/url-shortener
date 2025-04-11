package response

import (
	"time"

	"github.com/enesanbar/url-shortener/internal/domain"
)

type MappingPresenter struct {
}

func NewMappingPresenter() *MappingPresenter {
	return &MappingPresenter{}
}

func (mm *MappingPresenter) Single(mapping *domain.Mapping) *Response {
	var expiresAt *time.Time
	if mapping.ExpiresAt != nil {
		local := mapping.ExpiresAt.Local()
		expiresAt = &local
	} else {
		expiresAt = nil
	}

	m := &Response{
		ID:        mapping.ID,
		Code:      mapping.Code,
		URL:       mapping.URL,
		ExpiresAt: expiresAt,
		CreatedAt: mapping.CreatedAt.Local(),
		UpdatedAt: mapping.UpdatedAt.Local(),
	}

	return m
}

func (mm *MappingPresenter) Multiple(mappings []*domain.Mapping) []*Response {
	response := make([]*Response, 0, len(mappings))
	for _, mapping := range mappings {
		response = append(response, mm.Single(mapping))
	}
	return response
}
