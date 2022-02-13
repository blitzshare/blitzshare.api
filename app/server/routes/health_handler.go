package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary  health check
// @Schemes
// @Tags     healthcheck
// @Accept   text/plain
// @Produce  text/plain
// @Success  200  {string}  Im  alive
// @Router   /test [get]
func HealthHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		c.Writer.WriteString("im alive")
		c.Status(http.StatusOK)
	}
}
