package endpoints

import (
	"net/http"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	log "github.com/sirupsen/logrus"
)

func RegisterP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		var r model.P2pPeerRegistryCmd
		if err := c.ShouldBindWith(&r, binding.JSON); err == nil {
			log.Infoln("RegisterP2pPeer", r.Otp)

			msgId, err := deps.EventEmit.EmitP2pPeerRegistryCmd(deps.Config.Settings.QueueUrl, deps.Config.ClientId, &r)
			if err != nil {
				c.JSON(http.StatusInternalServerError, nil)
			} else {
				// TODO: return de-register token - to authenticate de-registration of otp
				c.JSON(http.StatusAccepted, gin.H{"ackId": msgId})
			}
		} else {
			log.Errorln("RegisterP2pPeer", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
