package endpoints

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"blitzshare.api/app/server/services/registry"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		var r model.PeerRegistry
		log.Infoln("RegisterP2pPeerHandler", r)
		if err := c.ShouldBindWith(&r, binding.JSON); err == nil {
			e := registry.RegisterPeer(deps, &r)
			if e != nil {
				log.Errorln("RegisterP2pPeerHandler", e)
				c.JSON(http.StatusInternalServerError, r)
			} else {
				c.JSON(http.StatusOK, r)
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	}
}
