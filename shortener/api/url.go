package api

import validation "github.com/go-ozzo/ozzo-validation"

type URL struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}

func (url *URL) ValidateShortenURL() error {
	return validation.ValidateStruct(&url,
		validation.Field(&url.Short, validation.Required))
}
func (url *URL) ValidateOriginURL() error {
	return validation.ValidateStruct(&url,
		validation.Field(&url.Original, validation.Required))
}
