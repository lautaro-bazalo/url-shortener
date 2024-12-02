package usecase

import (
	"go.uber.org/zap"
	"shortener/api"
	"shortener/internal/urlshort/model"
)

type Resposiotry interface {
	GetShortURL(url *model.URL) (*model.URL, error)
	CreateShortURL(url *model.URL) (*model.URL, error)
	DeleteShortURL(url *model.URL) error
	GetOriginalURL(url *model.URL) (*model.URL, error)
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

func (u *Usecase) GetShortURL(url api.URL) (*api.URL, error) {
	modelURL := ToModel(&url)

	if dbURL, err := u.repo.GetShortURL(modelURL); err != nil {
		return nil, err
	} else {
		return ToApi(dbURL), nil
	}

}
func (u *Usecase) CreateShortURL(url api.URL) (*api.URL, error) {
	modelURL := ToModel(&url)

	if dbURL, err := u.repo.CreateShortURL(modelURL); err != nil {
		return nil, err
	} else {
		return ToApi(dbURL), nil
	}

}
func (u *Usecase) DeleteShortURL(url api.URL) error {
	modelURL := ToModel(&url)
	return u.repo.DeleteShortURL(modelURL)
}
func (u *Usecase) GetOriginalURL(url api.URL) (*api.URL, error) {
	modelURL := ToModel(&url)

	if dbURL, err := u.repo.GetOriginalURL(modelURL); err != nil {
		return nil, err
	} else {
		return ToApi(dbURL), nil
	}
}
