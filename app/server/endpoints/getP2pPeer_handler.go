package endpoints

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/server/services/registry"
	"github.com/gin-gonic/gin"
)

func GetP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		log.Infoln("GetP2pPeerHandler", id)
		multiAddr, err := registry.GetPeerMultiAddr(deps, id)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"multiAddr": multiAddr})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		}
	}
}
