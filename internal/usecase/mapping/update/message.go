package update

import "time"

type MappingUpdatedPayload struct {
	ID        string     `json:"id" bson:"_id"`
	Code      string     `json:"code" bson:"code" example:"ZSDASZX"`
	URL       string     `json:"url" bson:"url" example:"https://example.com"`
	ExpiresAt *time.Time `json:"expires_at" bson:"expires_at" example:"2023-10-01T00:00:00Z"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at" example:"2023-10-01T00:00:00Z"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at" example:"2023-10-01T00:00:00Z"`
}
