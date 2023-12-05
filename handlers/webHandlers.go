package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type WebHandler struct{}

var date = time.Now()
var year = date.Year()

func (handler *WebHandler) IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK,
		"index.tmpl",
		gin.H{"title": "Home",
			"year": year,
		})
}
