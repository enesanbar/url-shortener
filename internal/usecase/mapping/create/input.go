package create

// Request data
type Request struct {
	Code      string `json:"code" example:"ZSDASZX" validate:"omitempty,gte=2,lte=15"`
	URL       string `json:"url" example:"https://example.com" validate:"required,is_url"`
	ExpiresAt string `json:"expires_at" example:"2006-01-02 15:04:05" validate:"omitempty,datetime=2006-01-02 15:04:05"`
}
