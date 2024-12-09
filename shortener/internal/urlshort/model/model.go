package model

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	ShortURL    string `gorm:"column:short_url"`
	OriginalURL string `gorm:"column:original_url;uniqueIndex"`
}

func (URL) TableName() string {
	return "url"
}
