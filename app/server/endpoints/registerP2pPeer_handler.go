package endpoints

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"blitzshare.api/app/server/services/registry"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		var r model.PeerRegistry
		if err := c.ShouldBindWith(&r, binding.JSON); err == nil {
			c.JSON(http.StatusOK, registry.RegisterPeer(deps, &r))
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		log.Infoln("RegisterP2pPeerHandler", r)
	}
}
