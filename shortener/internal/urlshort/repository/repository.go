package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hash/fnv"
	"shortener/internal/urlshort/model"
	"strings"
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

func (r *Resposiotry) GetShortURL(c context.Context, url *model.URL) (*model.URL, error) {

	value, err := r.cache.Get(c, url.OriginalURL).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			r.log.Warn("Not found in the cache", zap.String("url", url.OriginalURL), zap.Error(err))
		}
	} else {
		url.OriginalURL = value
		return url, nil
	}

	err = r.gorm.Transaction(func(tx *gorm.DB) error {
		return tx.Find(url).Error
	})

	if err != nil {
		r.gorm.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("I couldn't find a shortened url for: %s", url.ShortURL)
		}
		return nil, err
	}

	return url, nil
}

func (r *Resposiotry) CreateShortURL(c context.Context, url *model.URL) (*model.URL, error) {
	h := fnv.New32()
	_, err := h.Write([]byte(url.OriginalURL))

	if err != nil {
		return nil, err
	}
	hash := h.Sum(nil)

	url.ShortURL = fmt.Sprintf("%x", hash)

	err = r.gorm.Transaction(func(tx *gorm.DB) error {
		return tx.Create(url).Error
	})

	if err != nil {
		r.gorm.Rollback()
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, fmt.Errorf("already exist a url short for the url:  %s", url.OriginalURL)
		}
		return nil, fmt.Errorf("I couldn't storage the url %s: %s", url.OriginalURL, err)
	}
	if err := r.cache.Set(c, url.OriginalURL, url.ShortURL, 0).Err(); err != nil {
		r.log.Warn("couldn't store the original shortener", zap.String("url", url.OriginalURL), zap.Error(err))
	}

	if err = r.cache.Set(c, url.ShortURL, url.OriginalURL, 0).Err(); err != nil {
		r.log.Warn("couldn't store the shortener", zap.String("url", url.OriginalURL), zap.Error(err))
	}

	return url, nil
}

func (r *Resposiotry) DeleteShortURL(c context.Context, url *model.URL) error {
	_, err := r.cache.Del(c, url.OriginalURL).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			r.log.Warn("Not found in the cache for:", zap.String("url", url.OriginalURL), zap.Error(err))
		}
	}
	err = r.gorm.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(url).Error
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *Resposiotry) GetOriginalURL(c context.Context, url *model.URL) (*model.URL, error) {
	URLQuery := strings.Split(url.ShortURL, "/")

	url.ShortURL = URLQuery[len(URLQuery)-1]
	value, err := r.cache.Get(c, url.ShortURL).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			r.log.Warn("Not found in the cache for:", zap.String("url", url.ShortURL), zap.Error(err))
		}
	} else {
		r.log.Info("Found in the cache")
		url.OriginalURL = value
		return url, nil
	}

	err = r.gorm.Transaction(func(tx *gorm.DB) error {
		return tx.Find(url).Error
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("not found original URL for '%s' given", url.ShortURL)
		}
	}
	return url, nil
}
