package usecase

import (
	"context"
	"go.uber.org/zap"
	"shortener/api"
	"shortener/internal/urlshort/model"
)

type Resposiotry interface {
	GetShortURL(c context.Context, url *model.URL) (*model.URL, error)
	CreateShortURL(c context.Context, url *model.URL) (*model.URL, error)
	DeleteShortURL(c context.Context, url *model.URL) error
	GetOriginalURL(c context.Context, url *model.URL) (*model.URL, error)
}

type Usecase struct {
	repo Resposiotry
	log  *zap.Logger
}

func NewURLUsecase(repo Resposiotry, log *zap.Logger) *Usecase {
	return &Usecase{
		repo: repo,
		log:  log,
	}
}

func (u *Usecase) GetShortURL(c context.Context, url api.URL) (*api.URLResponse, error) {
	modelURL := ToModel(&url)

	if dbURL, err := u.repo.GetShortURL(c, modelURL); err != nil {
		return nil, err
	} else {
		return ToApi(dbURL), nil
	}

}
func (u *Usecase) CreateShortURL(c context.Context, url api.URL) (*api.URLResponse, error) {
	modelURL := ToModel(&url)

	if dbURL, err := u.repo.CreateShortURL(c, modelURL); err != nil {
		return nil, err
	} else {
		return ToApi(dbURL), nil
	}

}
func (u *Usecase) DeleteShortURL(c context.Context, url api.URL) error {
	modelURL := ToModel(&url)
	return u.repo.DeleteShortURL(c, modelURL)
}
func (u *Usecase) GetOriginalURL(c context.Context, url api.URL) (*api.URLResponse, error) {
	modelURL := ToModel(&url)

	if dbURL, err := u.repo.GetOriginalURL(c, modelURL); err != nil {
		return nil, err
	} else {
		return ToApi(dbURL), nil
	}
}
