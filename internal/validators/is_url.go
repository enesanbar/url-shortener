package validators

import (
	"net/url"
	"strings"

	"github.com/enesanbar/go-service/validation"

	"github.com/go-playground/validator/v10"
)

func NewIsURLValidator() validation.CustomValidation {
	return validation.CustomValidation{
		Tag:  "is_url",
		Func: IsURL,
		Messages: map[string]string{
			"en": "{0} must be a valid url",
			"tr": "{0} geçerli bir url olmalıdır",
		},
	}
}

// IsURL validates if a given string is a full url
func IsURL(fl validator.FieldLevel) bool {
	u, err := url.Parse(fl.Field().String())
	return err == nil && strings.HasPrefix(u.Scheme, "http") && u.Host != ""
}
