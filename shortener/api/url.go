package api

import validation "github.com/go-ozzo/ozzo-validation"

var MinValidationLengthValue = 22

type URL struct {
	RequestURL string `json:"request_url"`
}
type URLResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

func (url URL) Validate() error {
	return validation.ValidateStruct(&url,
		validation.Field(&url.RequestURL, validation.Required),
		validation.Field(&url.RequestURL, validation.Length(MinValidationLengthValue, 0)),
	)
}
