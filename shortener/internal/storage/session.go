package storage

import (
	"fmt"
	"shortener/internal/config"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Session struct {
	Log      *zap.Logger
	Database *gorm.DB
}

func NewSession(config config.Database, logger *zap.Logger) *Session {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sql, err := db.DB()

	if err != nil {
		panic(err)
	}

	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	sql.SetConnMaxLifetime(time.Hour)

	db.Exec("SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ")

	return &Session{
		Log:      logger,
		Database: db,
	}
}

func (s *Session) CloseDB() error {
	sqlDB, err := s.Database.DB()

	if err != nil {
		return err
	}

	err = sqlDB.Close()

	if err != nil {
		s.Log.Info("failed to close database connection", zap.Error(err))
		return err
	}

	return err
}
