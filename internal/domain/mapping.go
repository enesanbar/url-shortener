package domain

import (
	"time"

	"github.com/enesanbar/go-service/errors"
)

type Mapping struct {
	ID        string     `json:"id" bson:"_id"`
	Code      string     `json:"code" bson:"code" example:"ZSDASZX"`
	URL       string     `json:"url" bson:"url" example:"https://example.com"`
	ExpiresAt *time.Time `json:"expires_at" bson:"expires_at"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
}

func (m *Mapping) NewDateFromLayout(layout string, dateStr string) (*time.Time, error) {
	if dateStr != "" {
		location, err := time.LoadLocation("Europe/Istanbul")
		if err != nil {
			return nil, err
		}

		parsed, err := time.ParseInLocation(layout, dateStr, location)
		if err != nil {
			return nil, errors.Error{
				Code:    errors.EINVALID,
				Message: "cannot parse date",
				Err:     err,
			}
		}

		return &parsed, err
	}

	return nil, nil
}

func (m *Mapping) IsExpired() bool {
	if m.ExpiresAt != nil {
		return m.ExpiresAt.Before(time.Now())
	}

	return false
}
