package usecase

import (
	"shortener/api"
	"shortener/internal/urlshort/model"
)

const host = "https://me.li/"

func ToModel(url *api.URL) *model.URL {
	if len(url.RequestURL) == api.MinValidationLengthValue {
		return &model.URL{
			OriginalURL: url.RequestURL,
		}
	}

	return &model.URL{
		ShortURL: url.RequestURL,
	}
}

func ToApi(url *model.URL) *api.URLResponse {
	return &api.URLResponse{
		OriginalURL: url.OriginalURL,
		ShortURL:    host + url.ShortURL,
	}
}
