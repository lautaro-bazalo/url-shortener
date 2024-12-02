package urlshort

import (
	"github.com/gin-gonic/gin"
	"shortener/api"
)

type Usecase interface {
	CreateShortURL(apiURL api.URL) (*api.URL, error)
	GetShortURL(apiURL api.URL) (*api.URL, error)
	DeleteShortURL(apiURL api.URL) error
	GetOriginalURL(apiURL api.URL) (*api.URL, error)
}

type Handler struct {
	UseCase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{
		UseCase: usecase,
	}
}

func (h *Handler) AddHandler(r *gin.RouterGroup) {

	short := r.Group("/short")
	original := r.Group("/original")

	short.POST("", h.CreateShortURL)
	short.GET("", h.GetShortURL)
	short.DELETE("", h.DeleteShortURL)
	original.GET("", h.GetOriginalURL)
}

func (h *Handler) CreateShortURL(c *gin.Context) {}
func (h *Handler) GetShortURL(c *gin.Context)    {}
func (h *Handler) DeleteShortURL(c *gin.Context) {}
func (h *Handler) GetOriginalURL(c *gin.Context) {}
