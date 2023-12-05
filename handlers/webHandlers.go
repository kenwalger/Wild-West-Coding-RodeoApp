package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebHandler struct{}

func (handler *WebHandler) IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK,
		"index.tmpl",
		gin.H{"title": "Home"})
}
