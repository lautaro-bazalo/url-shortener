package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"log"
	"net/http"
	"shortener/internal/config"
	"shortener/internal/server/middleware"
	"time"
)

type Handler interface {
	AddHandler(r *gin.RouterGroup)
}

type Server struct {
	server   *http.Server
	log      *zap.Logger
	config   config.App
	handlers []Handler
}

func NewServer(log *zap.Logger, config config.App, handlers []Handler) *Server {
	return &Server{
		log:      log,
		config:   config,
		handlers: handlers,
	}

}

func (s *Server) Start() {

	newServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.newGinHandler(),
	}
	s.server = newServer

	go s.ListenAndServe()
}

func (s *Server) newGinHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()

	e.Use(middleware.PrometheusMiddleware())
	group := e.Group("/")

	for _, h := range s.handlers {
		h.AddHandler(group)
	}

	e.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return e
}

func (s *Server) ListenAndServe() {
	s.log.Info("starting server on port " + s.config.Port)
	if err := s.server.ListenAndServe(); err != nil {
		s.log.Warn("server stopped", zap.String("error", err.Error()))
	}

}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down the server: %v", err)
	}

}
