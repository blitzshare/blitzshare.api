package endpoints

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"blitzshare.api/app/server/services/events"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		var r model.PeerRegistry
		log.Infoln("RegisterP2pPeerHandler", r)
		if err := c.ShouldBindWith(&r, binding.JSON); err == nil {
			msgId, err := events.EmitP2pPeerRegistyEvent(deps, &r)
			if err != nil {
				log.Errorln("EmitP2pPeerRegistyEvent", err)
				c.JSON(http.StatusInternalServerError, gin.H{"ackId": msgId})
			} else {
				c.JSON(http.StatusAccepted, r)
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	}
}
