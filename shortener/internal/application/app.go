package application

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"os"
	"os/signal"
	cache "shortener/internal/cache"
	"shortener/internal/config"
	"shortener/internal/server"
	"shortener/internal/storage"
	"shortener/internal/urlshort"
	"shortener/internal/urlshort/repository"
	"shortener/internal/urlshort/usecase"
	"syscall"
)

type App struct {
	Session *storage.Session
	Server  *server.Server
	Logger  *zap.Logger
	Cache   *redis.Client
}

func NewApp() *App {
	config, err := config.ReadFromFile()

	if err != nil {
		panic(err)
	}
	logConfig := zap.NewProductionConfig()

	logger, err := logConfig.Build()

	if err != nil {
		panic(err)
	}

	session := storage.NewSession(config.Database, logger)

	cacheClient := cache.CacheClient(config.Cache, logger)

	repository := repository.NewRepository(session.Database, logger, cacheClient)

	urlUsecase := usecase.NewURLUsecase(repository, logger)

	handler := urlshort.NewHandler(urlUsecase)

	handlers := []server.Handler{handler}

	server := server.NewServer(logger, config.App, handlers)

	return &App{
		Session: session,
		Server:  server,
		Cache:   cacheClient,
	}

}

func (a *App) Run() {
	defer a.Server.Shutdown()

	a.Server.Start()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := a.Session.CloseDB(); err != nil {
		a.Logger.Error("failed to close session", zap.Error(err))
	}

	if err := a.Cache.Close(); err != nil {
		a.Logger.Error("failed to close cache", zap.Error(err))
	}

}
