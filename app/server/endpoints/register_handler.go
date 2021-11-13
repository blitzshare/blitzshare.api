package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Writer.WriteString("im alive")
		c.Status(http.StatusOK)
	}
}
