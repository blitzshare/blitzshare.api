package routes

import (
	"blitzshare.api/app/services/key"
	"github.com/gin-gonic/gin"
)

const (
	ServiceHeader = "X-Blitzshare-Service"
	KeyHeader     = "X-Api-Key"
	ServiceName   = "blitzshare.api.v1"
)

func AddDefaultResponseHeaders(c *gin.Context) {
	c.Header(ServiceHeader, ServiceName)
}

func GetApiKeyHeader(c *gin.Context) *string {
	apiKeyHeader := c.Request.Header.Get(KeyHeader)
	return &apiKeyHeader
}

func IsNotAuthorized(c *gin.Context, chain key.ApiKeychain) bool {
	key := GetApiKeyHeader(c)
	return chain.IsValid(key) == false
}
