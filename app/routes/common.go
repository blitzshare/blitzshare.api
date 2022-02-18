package routes

import "github.com/gin-gonic/gin"

const (
	ServiceHeader = "X-Blitzshare-Service"
)

func AddDefaultResponseHeaders(c *gin.Context) {
	c.Header(ServiceHeader, "blitzshare.api")
}

func GetApiKeyHeader(c *gin.Context) *string {
	apiKeyHeader := c.Request.Header.Get("x-api-key")
	return &apiKeyHeader
}
