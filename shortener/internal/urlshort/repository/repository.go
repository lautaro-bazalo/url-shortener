package repository

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"shortener/internal/urlshort/model"
)

type Resposiotry struct {
	gorm  *gorm.DB
	log   *zap.Logger
	cache *redis.Client
}

func NewRepository(db *gorm.DB, logger *zap.Logger, cache *redis.Client) *Resposiotry {
	return &Resposiotry{
		gorm:  db,
		log:   logger,
		cache: cache,
	}
}

func (r *Resposiotry) GetShortURL(url *model.URL) (*model.URL, error) {
	err := r.gorm.Transaction(func(tx *gorm.DB) error {
		return tx.Find(url).Error
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		}
	}

	return url, nil
}

func (r *Resposiotry) CreateShortURL(url *model.URL) (*model.URL, error) {

	err := r.gorm.Transaction(func(tx *gorm.DB) error {
		return tx.Create(url).Error
	})
	if err != nil {

	}

	return url, nil
}

func (r *Resposiotry) DeleteShortURL(url *model.URL) error {
	err := r.gorm.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(url).Error
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Resposiotry) GetOriginalURL(url *model.URL) (*model.URL, error) {
	err := r.gorm.Transaction(func(tx *gorm.DB) error {
		return tx.Find(url).Error
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("not found original URL for '%s' given", url.ShortURL)
		}
	}
	return url, nil
}
