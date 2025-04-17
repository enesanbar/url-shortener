package domain

type MappingCreatedPayload struct {
	MappingID   string `json:"mappingId"`
	ShortURL    string `json:"shortUrl"`
	OriginalURL string `json:"originalUrl"`
}
