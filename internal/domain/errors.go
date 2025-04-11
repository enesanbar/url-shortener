package domain

import "errors"

var (
	// ErrMappingNotFound is returned when a mapping lookup fails.
	ErrMappingNotFound = errors.New("mapping not found")
)
