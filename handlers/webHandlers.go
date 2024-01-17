package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type WebHandler struct{}

// Package-level variables
var date = time.Now()
var year = date.Year()

// swagger:operation GET / web index
// The application index page
// ---
// produces:
// - application/html
// responses:
//
//	'200':
//	    description: Successful rendering of the index page
func (handler *WebHandler) IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Home"})
}
