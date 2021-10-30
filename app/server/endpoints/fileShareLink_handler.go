package endpoints

import (
	"net/http"

	dep "blitzshare.fileshare.api/app/dependencies"
	"blitzshare.fileshare.api/app/server/services"
	"github.com/gin-gonic/gin"
)

func FileShareHandler(deps *dep.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		services.SubmitUploadMsgEvent(deps.Config.Settings.QueueUrl)
		presignedUrl := services.GetPresignedUrl(deps)

		c.JSON(http.StatusOK, gin.H{
			"uploadUrl":    presignedUrl.Url,
			"expirationMs": presignedUrl.ExpirationMs,
		})
		c.Status(http.StatusOK)
	}
}
