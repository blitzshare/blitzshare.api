package routes

import (
	"blitzshare.api/app/services/key"
	"github.com/gin-gonic/gin"
)

const (
	ServiceHeader = "X-Blitzshare-Service"
)

func AddDefaultResponseHeaders(c *gin.Context) {
	c.Header(ServiceHeader, "blitzshare.api")
}

func GetApiKeyHeader(c *gin.Context) *string {
	apiKeyHeader := c.Request.Header.Get("X-Api-Key")
	return &apiKeyHeader
}

func IsNotAuthorized(c *gin.Context, chain key.ApiKeychain) bool {
	key := GetApiKeyHeader(c)
	return chain.IsValid(key) == false

}
