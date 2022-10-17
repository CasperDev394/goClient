package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	c.R.GET("/bot", h.Bot)

	g := c.R.Group("/bot")

	g.GET("/ping", h.Ping)

}

func (h *Handler) Bot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}
