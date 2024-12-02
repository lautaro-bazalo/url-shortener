package usecase

import (
	"shortener/api"
	"shortener/internal/urlshort/model"
)

func ToApi(url *model.URL) *api.URL {
	return &api.URL{
		Original: url.OriginalURL,
		Short:    url.ShortURL,
	}
}

func ToModel(url *api.URL) *model.URL {
	return &model.URL{
		OriginalURL: url.Original,
		ShortURL:    url.Short,
	}
}
