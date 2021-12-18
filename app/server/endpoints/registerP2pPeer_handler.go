package endpoints

import (
	"net/http"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"blitzshare.api/app/server/services/events"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		var r model.P2pPeerRegistryCmd
		if err := c.ShouldBindWith(&r, binding.JSON); err == nil {
			msgId, err := events.EmitP2pPeerRegistryCmd(deps, &r)
			if err != nil {
				c.JSON(http.StatusInternalServerError, nil)
			} else {
				// TODO: return de-register token - to authenticate de-registration of otp
				c.JSON(http.StatusAccepted, gin.H{"ackId": msgId})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
