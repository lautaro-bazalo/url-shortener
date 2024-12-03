package urlshort

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *Handler) CreateShortURL(c *gin.Context) {
	apiURL := &api.URL{}
	if err := c.Bind(apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error unmarshalling request": err.Error()})
	}

	if err := apiURL.ValidateShortenURL(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation shorten error": err.Error()})
	}

	if apiURL, err := h.UseCase.CreateShortURL(*apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, apiURL)
	}
}
func (h *Handler) GetShortURL(c *gin.Context) {
	apiURL := &api.URL{}
	if err := c.Bind(apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error unmarshalling request": err.Error()})
	}

	if err := apiURL.ValidateOriginURL(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation original error": err.Error()})
	}

	if apiURL, err := h.UseCase.GetShortURL(*apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, apiURL)
	}
}
func (h *Handler) DeleteShortURL(c *gin.Context) {
	apiURL := &api.URL{}
	if err := c.Bind(apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error unmarshalling request": err.Error()})
	}
	if err := apiURL.ValidateShortenURL(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation shorten error": err.Error()})
	}
	if err := h.UseCase.DeleteShortURL(*apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusNoContent)
	}

}
func (h *Handler) GetOriginalURL(c *gin.Context) {
	url := &api.URL{}
	if err := c.Bind(url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error unmarshalling request": err.Error()})
	}

	if err := url.ValidateShortenURL(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation shorten error": err.Error()})
	}
}
