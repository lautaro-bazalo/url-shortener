package urlshort

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"shortener/api"
)

type Usecase interface {
	CreateShortURL(c context.Context, apiURL api.URL) (*api.URLResponse, error)
	GetShortURL(c context.Context, apiURL api.URL) (*api.URLResponse, error)
	DeleteShortURL(c context.Context, apiURL api.URL) error
	GetOriginalURL(c context.Context, apiURL api.URL) (*api.URLResponse, error)
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

	root := r.Group("/")

	root.POST("", h.CreateShortURL)
	root.GET("", h.GetURL)
	root.DELETE("", h.DeleteShortURL)
}

func (h *Handler) CreateShortURL(c *gin.Context) {
	apiURL := &api.URL{}
	if err := c.BindJSON(apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error unmarshalling request": err.Error()})
	}

	if err := apiURL.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}
	ctx := context.Background()

	if apiURL, err := h.UseCase.CreateShortURL(ctx, *apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, apiURL)
	}

}
func (h *Handler) GetURL(c *gin.Context) {
	apiURL := &api.URL{}
	if err := c.BindJSON(apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error unmarshalling request": err.Error()})
	}

	if err := apiURL.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation original error": err.Error()})
		return
	}

	ctx := context.Background()

	h.GETHandler(c, apiURL, ctx)

}

func (h *Handler) DeleteShortURL(c *gin.Context) {
	apiURL := &api.URL{}
	if err := c.Bind(apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error unmarshalling request": err.Error()})
	}
	if err := apiURL.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation shorten error": err.Error()})
		return
	}
	ctx := context.Background()

	if err := h.UseCase.DeleteShortURL(ctx, *apiURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusNoContent)
	}

}

func (h *Handler) GETHandler(c *gin.Context, apiURL *api.URL, ctx context.Context) {
	if len(apiURL.RequestURL) > api.MinValidationLengthValue {
		if apiURL, err := h.UseCase.GetShortURL(ctx, *apiURL); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, apiURL)
		}
	} else {
		if apiURL, err := h.UseCase.GetOriginalURL(ctx, *apiURL); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, apiURL)
		}
	}
}
