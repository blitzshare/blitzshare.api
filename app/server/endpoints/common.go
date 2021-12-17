package endpoints

import "github.com/gin-gonic/gin"

const (
	ServiceHeader = "X-Blitzshare-Service"
)

func AddDefaultResponseHeaders(c *gin.Context) {
	c.Header(ServiceHeader, "blitzshare.api")
}
