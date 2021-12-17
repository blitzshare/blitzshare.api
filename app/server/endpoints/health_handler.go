package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		c.Writer.WriteString("im alive")
		c.Status(http.StatusOK)
	}
}
