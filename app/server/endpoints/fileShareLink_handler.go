package endpoints

import (
	"net/http"

	dep "blitzshare.api/app/dependencies"
	"blitzshare.api/app/server/services"
	"github.com/gin-gonic/gin"
)

func FileShareHandler(deps *dep.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		presignedUrl := services.GetPresignedUrl(deps)

		c.JSON(http.StatusOK, gin.H{
			"uploadUrl":    presignedUrl.Url,
			"expirationMs": presignedUrl.ExpirationMs,
		})
		c.Status(http.StatusOK)
	}
}
